package middleware

import (
	"net/http"
	"os"
	"sample-server/internal/config"

	"github.com/labstack/echo/v4"
)

type APIKeyAuthConfig struct {
	APIKeys map[string]struct{}
	IsLocal bool
}

type APIKeyAuthMiddleware struct {
	conf *APIKeyAuthConfig
}

func NewAPIKeyAuthMiddleware(conf *APIKeyAuthConfig) *APIKeyAuthMiddleware {
	return &APIKeyAuthMiddleware{conf: conf}
}

func (m *APIKeyAuthMiddleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-KEY")
		if _, found := m.conf.APIKeys[apiKey]; found {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or missing API key"})
	}
}

// 設定例: 環境変数API_KEYS="key1,key2" で複数APIキーを許可
func LoadAPIKeyAuthConfig() *APIKeyAuthConfig {
	apiKeys := make(map[string]struct{})
	for _, k := range splitAndTrim(os.Getenv("API_KEYS")) {
		if k != "" {
			apiKeys[k] = struct{}{}
		}
	}
	isLocal := os.Getenv("ENV") == "local"
	return &APIKeyAuthConfig{APIKeys: apiKeys, IsLocal: isLocal}
}

// ConfigからAPIキーを受け取る形に修正
func LoadAPIKeyAuthConfigFromConfig(cfg *config.Config) *APIKeyAuthConfig {
	apiKeys := make(map[string]struct{})
	for k := range cfg.APIKeys {
		apiKeys[k] = struct{}{}
	}
	return &APIKeyAuthConfig{APIKeys: apiKeys, IsLocal: false}
}

func splitAndTrim(s string) []string {
	var res []string
	for _, v := range splitComma(s) {
		res = append(res, trimSpace(v))
	}
	return res
}

func splitComma(s string) []string {
	var r []string
	for _, v := range []rune(s) {
		if v == ',' {
			r = append(r, "")
		} else {
			if len(r) == 0 {
				r = append(r, "")
			}
			r[len(r)-1] += string(v)
		}
	}
	return r
}
func trimSpace(s string) string { return s }
