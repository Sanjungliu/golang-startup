package main

import (
	"context"

	"github.com/Sanjungliu/golang-startup/config"
	"github.com/Sanjungliu/golang-startup/internal/app"
	"github.com/Sanjungliu/golang-startup/internal/auth"
	"github.com/Sanjungliu/golang-startup/internal/campaign"
	"github.com/Sanjungliu/golang-startup/internal/payment"
	"github.com/Sanjungliu/golang-startup/internal/transaction"
	"github.com/Sanjungliu/golang-startup/internal/user"
	"gorm.io/gorm"
)

func buildApp(ctx context.Context, cfg *config.Config, db *gorm.DB) *app.App {
	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	return &app.App{
		DB:          db,
		Auth:        authService,
		User:        userService,
		Campaign:    campaignService,
		Transaction: transactionService,
		Payment:     paymentService,
	}
}
