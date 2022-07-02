package fee_schedule_server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	qualdevlabs_auth_go_client "github.com/Jriles/QualDevLabsAuthGoClient"
	"github.com/gin-gonic/gin"
)

// ApiMiddleware will add the db connection to the context
func DBMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}

func AuthMiddleWare(c *gin.Context) {
	orgId := os.Getenv("AUTH_ORG_ID") // string | the org's UUID (unique)
	appId := os.Getenv("AUTH_APP_ID") // string | the app's UUID (unique)
	authApiKey := os.Getenv("AUTH_API_KEY")
	authApiKeyStruct := qualdevlabs_auth_go_client.APIKey{
		Key: authApiKey,
	}

	sessionToken := c.Request.Header["Session_token"][0]
	authTokenHeaderStruct := qualdevlabs_auth_go_client.APIKey{
		Key: sessionToken,
	}

	userId := c.Request.Header["User_id"][0]

	configuration := qualdevlabs_auth_go_client.NewConfiguration()
	api_client := qualdevlabs_auth_go_client.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), qualdevlabs_auth_go_client.ContextAPIKeys, map[string]qualdevlabs_auth_go_client.APIKey{
		"apiKeyHeader": authApiKeyStruct,
		"tokenHeader":  authTokenHeaderStruct,
	})
	resp, err := api_client.DefaultApi.ValidateSession(ctx, orgId, appId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ValidateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	if resp.StatusCode == 200 {
		c.Next()
	}
}
