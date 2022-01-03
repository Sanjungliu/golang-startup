package httpserver

import (
	"net/http"
	"strings"

	"github.com/Sanjungliu/golang-startup/internal/app"
	"github.com/Sanjungliu/golang-startup/internal/auth"
	"github.com/Sanjungliu/golang-startup/internal/httpserver/handler"
	"github.com/Sanjungliu/golang-startup/internal/user"
	"github.com/Sanjungliu/golang-startup/pkg/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(app *app.App) *http.Server {
	s := &http.Server{
		Handler: buildRoute(app),
	}
	return s
}

func buildRoute(app *app.App) http.Handler {
	userHandler := handler.NewUserHandler(app.User, app.Auth)
	campaignHandler := handler.NewCampaignHandler(app.Campaign)
	transactionHandler := handler.NewTransactionHandler(app.Transaction)

	auth := authMiddleware(app.Auth, app.User)

	router := gin.Default()
	router.Use(
		cors.Default(),
	)

	// router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/sessions", userHandler.Login)
	api.POST("/users", userHandler.RegisterUser)

	api.Use(auth).GET("/users", userHandler.FetchUser)
	api.Use(auth).POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.Use(auth).POST("/avatars", userHandler.UploadAvatar)

	api.Use(auth).GET("/campaigns", campaignHandler.GetCampaigns)
	api.Use(auth).GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.Use(auth).POST("/campaigns", campaignHandler.CreateCampaign)
	api.Use(auth).PUT("/campaigns/:id", campaignHandler.UpdateCampaign)
	api.Use(auth).POST("/campaign-images", campaignHandler.UploadImage)

	api.Use(auth).GET("campaigns/:id/transactions", transactionHandler.GetCampaignTransactions)
	api.Use(auth).GET("transactions", transactionHandler.GetUserTransactions)
	api.Use(auth).POST("transactions", transactionHandler.CreateTransaction)
	api.Use(auth).POST("transactions/notifications", transactionHandler.GetNotification)

	router.Run()
	return router
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
