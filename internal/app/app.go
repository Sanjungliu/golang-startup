package app

import (
	"github.com/Sanjungliu/golang-startup/internal/auth"
	"github.com/Sanjungliu/golang-startup/internal/campaign"
	"github.com/Sanjungliu/golang-startup/internal/payment"
	"github.com/Sanjungliu/golang-startup/internal/transaction"
	"github.com/Sanjungliu/golang-startup/internal/user"
	"gorm.io/gorm"
)

type App struct {
	DB          *gorm.DB
	Auth        auth.Service
	User        user.Service
	Campaign    campaign.Service
	Transaction transaction.Service
	Payment     payment.Service
}
