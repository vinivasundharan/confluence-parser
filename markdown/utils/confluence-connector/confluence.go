package confluence

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	HTTPHelper "markdown/markdown/utils/http"
	"net/http"

	"github.com/caarlos0/env"
)

type confAuthentication struct {
	UserName string `env:"CONFLUENCE_USERNAME" envDefault:"vinivasundharan2011@gmail.com"`
	URL      string `env:"CONFLUENCE_URL" envDefault:"https://vinivasundharan.atlassian.net"`
	Password string `env:"CONFLUENCE_PASSWORD" envDefault:"ATATT3xFfGF0NPVT32-8i0-dONLMQTaqSCH-Lr7gC6Ew04iwzfdbC-MIB4cATjeGeCxGKm1_5KpsRwGU6ccVb4QiVthhLa4Gn7J55ht9ekInk1LJm0mogigIigVN-9tQrPg_t0RB6tU6gI9lKdISKXe1TMF5eNzTxt1QVv_gRDwgb6CFqvEj8N0=E60B440A"`
}

const (
	ENDPOINT_API_VERSION = "wiki/rest/api"
	ENDPOINT_CONTENT     = "content"
	ENDPOINT_SPACE       = "space"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("test", "test"))
	return nil
}

func confRequest(reqMethod string, reqURL string) (resp *http.Response) {
	confClient := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	request, _ := http.NewRequest(reqMethod, reqURL, nil)
	auth := confAuthentication{}
	parseerror := env.Parse(&auth)
	if parseerror != nil {
		log.Fatalf("unable to parse environment variables: %e", parseerror)
		panic(parseerror)
	}
	request.Header.Add("Authorization", "Basic "+basicAuth(auth.UserName, auth.Password))
	request.Header.Add("Content-Type", "application/json")
	resp, _ = confClient.Do(request)
	return resp
}

func ConfPageContent(contentID string) (responseJSON []byte) {
	httpCodes := HTTPHelper.HTTPCodes{}
	auth := confAuthentication{}
	parseerror := env.Parse(&auth)
	if parseerror != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", parseerror)
		panic(parseerror)
	}
	var contentURL = auth.URL + "/" + ENDPOINT_API_VERSION + "/" + ENDPOINT_CONTENT + "/" + contentID + "?expand=body.storage"
	resp := confRequest(httpCodes.GET, contentURL)
	if resp.StatusCode > 299 {
		log.Fatalf("\nRequest was not successful. The request returned %s", resp.Status)
		return
	}
	responseJSON, _ = ioutil.ReadAll(resp.Body)
	return responseJSON
}

func confAuth() (confAuth *confAuthentication) {
	confAuth = &confAuthentication{}
	return confAuth
}

func confUserName() (username string) {
	auth := confAuthentication{}
	parseerror := env.Parse(&auth)
	if parseerror != nil {
		log.Fatalf("unable to parse environment variables: %e", parseerror)
		panic(parseerror)
	}
	return auth.UserName
}

func confPassword() (username string) {
	auth := confAuthentication{}
	parseerror := env.Parse(&auth)
	if parseerror != nil {
		log.Fatalf("unable to parse environment variables: %e", parseerror)
		panic(parseerror)
	}
	return auth.Password
}

func BuildSpaceURL(space, content string) (URL string) {
	return "g"
}

func BuildContentURL(content string) (URL string) {

	return "g"
}

func getSpaceID(content Content) (spaceID int) {
	if content.Type != "page" {
		log.Fatal("Currently the space ID can be retrieved for confluence content of type `page`")
		return
	}
	auth := confAuthentication{}
	parseError := env.Parse(&auth)
	fmt.Print(parseError)
	return 0
}

func getContentID(content string) (contentID int) {
	return 0
}
