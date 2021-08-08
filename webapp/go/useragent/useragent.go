package useragent

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(?i)(bot|crawler|spider)(?:[-_ .\/;@()]|$)`)

func IsBot(ua string) bool {
	// /ISUCONbot(-Mobile)?/
	if strings.Contains(ua, "ISUCONbot") || strings.Contains(ua, "ISUCONbot-Mobile") {
		return true
	}

	// /ISUCONbot-Image\//
	if strings.Contains(ua, "ISUCONbot-Image/") {
		return true
	}

	// /Mediapartners-ISUCON/
	if strings.Contains(ua, "Mediapartners-ISUCON") {
		return true
	}

	// /ISUCONCoffee/
	if strings.Contains(ua, "ISUCONCoffee") {
		return true
	}

	// /ISUCONFeedSeeker(Beta)?/
	if strings.Contains(ua, "ISUCONFeedSeeker") || strings.Contains(ua, "ISUCONFeedSeekerBeta") {
		return true
	}

	// /crawler \(https:\/\/isucon\.invalid\/(support\/faq\/|help\/jp\/)/
	if strings.Contains(ua, "crawler (https://isucon.invalid/support/faq/") || strings.Contains(ua, "crawler (https://isucon.invalid/help/jp/") {
		return true
	}

	// /isubot/
	if strings.Contains(ua, "isubot") {
		return true
	}

	// /Isupider/
	if strings.Contains(ua, "Isupider") {
		return true
	}

	// /Isupider(-image)?\+/
	if strings.Contains(ua, "Isupider+") || strings.Contains(ua, "Isupider-image+") {
		return true
	}

	// /(bot|crawler|spider)(?:[-_ .\/;@()]|$)/i
	if re.MatchString(ua) {
		return true
	}

	return false
}
