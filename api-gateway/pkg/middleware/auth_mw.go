package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	authpb "github.com/kirpepa/spendly-api/auth/proto"
)

type AuthMiddleware struct {
	AuthClient authpb.AuthServiceClient
	Timeout    time.Duration
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing authorization header"})
			return
		}
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization header"})
			return
		}
		token := parts[1]

		ctx, cancel := context.WithTimeout(context.Background(), m.Timeout)
		defer cancel()

		resp, err := m.AuthClient.ValidateToken(ctx, &authpb.ValidateRequest{Token: token})
		if err != nil || !resp.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user_id", resp.UserId)
		c.Set("user_email", resp.Email)
		c.Next()
	}
}
