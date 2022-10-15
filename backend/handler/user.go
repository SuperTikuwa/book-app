package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/SuperTikuwa/book_app/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type StoreUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	body := StoreUserRequest{}
	if err := json.Unmarshal(b, &body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		},
	))

	client := cognitoidentityprovider.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

	res, err := client.SignUp(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("AWS_COGNITO_CLIENT_ID")),
		Username: aws.String(body.Name),
		Password: aws.String("Password@123"),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(body.Email),
			},
		},
		SecretHash: aws.String(generateSecretHash(body.Name)),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	u := model.User{
		Name:        body.Name,
		Email:       body.Email,
		CognitoUUID: *res.UserSub,
	}

	if err := u.Create(); err != nil {
		log.Println("Cognitoからユーザーを削除します")
		client.AdminDeleteUser(&cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: aws.String(os.Getenv("AWS_COGNITO_USER_POOL_ID")),
			Username:   aws.String(body.Name),
		})

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
