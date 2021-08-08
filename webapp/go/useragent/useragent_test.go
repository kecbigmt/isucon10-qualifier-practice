package useragent

import (
	"testing"
)

func TestCheckIfBot(t *testing.T) {
	botPhrases := []string{
		"ISUCONbot",
		"ISUCONbot-Mobile",
		"ISUCONbot-Image/",
		"Mediapartners-ISUCON",
		"ISUCONCoffee",
		"ISUCONFeedSeeker",
		"crawler (https://isucon.invalid/support/faq/",
		"crawler (https://isucon.invalid/help/jp/",
		"isubot",
		"Isupider",
		"Isupider+",
		"Isupider-image+",
		"bot-",
		"bot_",
		"bot ",
		"bot.",
		"bot/",
		"bot;",
		"bot@",
		"bot(",
		"bot)",
		"crawler-",
		"crawler_",
		"crawler ",
		"crawler.",
		"crawler/",
		"crawler;",
		"crawler@",
		"crawler(",
		"crawler)",
		"spider-",
		"spider_",
		"spider ",
		"spider.",
		"spider/",
		"spider;",
		"spider@",
		"spider(",
		"spider)",
		"Bot-",
		"Bot_",
		"Bot ",
		"Bot.",
		"Bot/",
		"Bot;",
		"Bot@",
		"Bot(",
		"Bot)",
		"Crawler-",
		"Crawler_",
		"Crawler ",
		"Crawler.",
		"Crawler/",
		"Crawler;",
		"Crawler@",
		"Crawler(",
		"Crawler)",
		"Spider-",
		"Spider_",
		"Spider ",
		"Spider.",
		"Spider/",
		"Spider;",
		"Spider@",
		"Spider(",
		"Spider)",
	}

	botPatterns := []string{
		"---bot",
		"---Bot",
		"---crawler",
		"---Crawler",
		"---spider",
		"---Spider",
	}
	for _, phrase := range botPhrases {
		botPatterns = append(botPatterns, phrase)
		botPatterns = append(botPatterns, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) "+phrase+" Chrome/63.0.3239.132 Safari/537.36")
	}

	for _, pattern := range botPatterns {
		if isBot := IsBot(pattern); !isBot {
			t.Fatalf("%s is bot\n", pattern)
		}
	}

	notBotPatterns := []string{
		"bottom",
		"spiderman",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) Chrome/63.0.3239.132 Safari/537.36",
	}

	for _, pattern := range notBotPatterns {
		if isBot := IsBot((pattern)); isBot {
			t.Fatalf("%s is not bot\n", pattern)
		}
	}
}
