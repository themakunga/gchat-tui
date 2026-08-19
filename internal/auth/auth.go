package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gchat-tui/internal/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/chat/v1"
)

func GetClient(ctx context.Context, paths config.AppPaths) *http.Client {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Make sure you have 'credentials.json' in the '~/gchat-tui' directory")
	}

	conf, err := google.ConfigFromJSON(b, chat.ChatReadonlyScope, "https://www.googleapis.com/auth/chat.messages")
	if err != nil {
		log.Fatalf("Error parsing credentials.json: %v", err)
	}

	tokFile := filepath.Join(paths.AuthDir, "token.json")
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(ctx, conf)
		saveToken(tokFile, tok)
	}

	return conf.Client(ctx, tok)
}

func getTokernFromWeb(ctx context.Context, conf *oauth2.Config) *oauth2.Token {

}
