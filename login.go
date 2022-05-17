package adpexpert

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

func (c *Client) Login(username, password string) error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatalf("failed to create cookie jar: %+v", err)
	}

	client := resty.New().
		SetRedirectPolicy(resty.RedirectPolicyFunc(func(r1 *http.Request, r2 []*http.Request) error {
			return http.ErrUseLastResponse
		})).
		SetCookieJar(jar).
		SetDebug(true)

	target := "https://expert.brasil.adp.com/expert/"
	resp, err := client.R().
		SetFormData(map[string]string{
			"USER":     username,
			"PASSWORD": password,
			"TARGET":   target,
		}).
		Post(c.getLoginURL())
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatalf("login failed: %+v", err)
	}
	if resp.StatusCode() != 302 {
		fmt.Println("Status:", resp.Status)
		log.Fatalf("login failed with invalid status: %s", resp.Status())
	}
	loc, err := resp.RawResponse.Location()
	if loc == nil || loc.String() != target {
		log.Fatalf("login failed with invalid redirect location: %s", loc)
	}
	// client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(3))
	resp, err = client.R().
		SetDoNotParseResponse(true).
		Get("https://expert.brasil.adp.com/expert/v4/?lp=true")
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatalf("load session ID failed: %+v", err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp.RawResponse)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatalf("load session ID failed: %+v", err)
	}

	sessionID, exists := doc.Find("input#newexpert_sessionid").First().Attr("value")
	if !exists {
		log.Fatalf("load session ID failed because session ID was not in the page")
	}

	fmt.Println("SessionID:", sessionID)

	resp, err = client.R().
		SetHeader("newexpert_sessionid", sessionID).
		SetBody(map[string]any{
			"punchType":      "SPDesktop",
			"punchLatitude":  nil,
			"punchLongitude": nil,
			"punchAction":    nil,
		}).
		Get("https://expert.brasil.adp.com/expert/api/punch/punchin?lp=true")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	log.Fatalf("punch-in failed: %+v", err)
	// }
	// if !resp.IsSuccess() {
	// 	log.Fatalf("punch-in failed with invalid status: %s", resp.Status())
	// }

	return err
}
