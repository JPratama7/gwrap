package gwrap

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

func GetClient(config *oauth2.Config, fileName string) *http.Client {
	tok, err := TokenFromFile(fileName)
	if err != nil {
		tok = getTokenFromWeb(config)
		SaveToken(fileName, tok)
	}

	tok, err = RefreshToken(context.Background(), config, tok).Token()

	if err != nil {
		return nil
	}
	SaveToken(fileName, tok)

	return config.Client(context.Background(), tok)
}
