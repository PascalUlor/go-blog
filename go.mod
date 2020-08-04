module example.com/main

go 1.14

require (
	example.com/me/api v0.0.0
	example.com/me/auth v0.0.0 // indirect
	example.com/me/controllers v0.0.0
	example.com/me/formaterror v0.0.0 // indirect
	example.com/me/middlewares v0.0.0 // indirect
	example.com/me/models v0.0.0
	example.com/me/responses v0.0.0 // indirect
	example.com/me/seed v0.0.0 // indirect
	github.com/badoux/checkmail v0.0.0-20181210160741-9661bd69e9ad // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20200422194213-44a606286825 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1
)

replace example.com/me/api => ./api

replace example.com/me/controllers => ./api/controllers

replace example.com/me/seed => ./api/seed

replace example.com/me/responses => ./api/responses

replace example.com/me/auth => ./api/auth

replace example.com/me/models => ./api/models

replace example.com/me/formaterror => ./api/utils/formaterror

replace example.com/me/middlewares => ./api/middlewares
