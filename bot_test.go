package idobot_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mk2/idobot"
)

func isValidBot(t *testing.T, bot idobot.Bot, err error) {
	t.Helper()
	if bot == nil || err != nil {
		t.Fatalf("botが正しく生成されていません。\n")
	}
}

func TestIdobot_新しくbotを生成できるか(t *testing.T) {
	url := "url"
	apiToken := "token"
	userAgent := "userAgent"
	onStart := func(_ idobot.Bot, _ *idobot.SeedMsg) {}
	onEvent := func(_ idobot.Bot, _ *idobot.EventMsg) {}
	onError := func(_ idobot.Bot, _ error) {}
	bot, err := idobot.NewBot(&idobot.NewBotOpts{
		URL:           url,
		APIToken:      apiToken,
		UserAgent:     userAgent,
		StoreFilePath: "./test.db",
		OnStart:       onStart,
		OnEvent:       onEvent,
		OnError:       onError,
	})
	defer bot.Stop()
	defer os.Remove("./test.db")

	isValidBot(t, bot, err)
}

func TestIdobot_PostMessage実行がうまくいくかどうか(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("body"))
	}))
	url := testServer.URL
	apiToken := "token"
	userAgent := "userAgent"
	onStart := func(_ idobot.Bot, _ *idobot.SeedMsg) {}
	onEvent := func(_ idobot.Bot, _ *idobot.EventMsg) {}
	onError := func(_ idobot.Bot, _ error) {}
	bot, err := idobot.NewBot(&idobot.NewBotOpts{
		URL:           url,
		APIToken:      apiToken,
		UserAgent:     userAgent,
		StoreFilePath: "./test.db",
		OnStart:       onStart,
		OnEvent:       onEvent,
		OnError:       onError,
	})
	defer bot.Stop()
	defer os.Remove("./test.db")

	isValidBot(t, bot, err)

	body, err := bot.PostMessage(100, "test")

	if body != "body" {
		t.Errorf("PostMessage respone: expected: \"body\"\n received: \"%s\"", body)
	}

	if err != nil {
		t.Errorf("PostMessage returns error")
	}
}
