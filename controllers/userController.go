package controllers

import (
	"github.com/Golang-Authentication/database"
	"github.com/Golang-Authentication/vendor/github.com/go-playground/validator/v10"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func SignUp()
func Login()
func GetUsers()
func GetUser()
