package controllers

import (
	"Go-JWT/helpers"
	helper "Go-JWT/helpers"
	"Go_JWT/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PraveenAththanayake/Go-JWT/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10" 
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crpto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

func SignUp() {

}

func Login() {

}

func GetUsers() {

}

func GetUser() {

}