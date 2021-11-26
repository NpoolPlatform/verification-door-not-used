package googlerecaptcha

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/verification-door/message/npool"
	"golang.org/x/xerrors"
)

const (
	KeyRecaptchaURL    = "recaptcha_url"
	KeyRecaptchaSecret = "recaptcha_secret"
)

func VerifyGoogleRecaptcha(in *npool.VerifyGoogleRecaptchaRequest) (*npool.VerifyGoogleRecaptchaResponse, error) {
	hostname := config.GetStringValueWithNameSpace("", config.KeyHostname)
	recaptchaURL := config.GetStringValueWithNameSpace(hostname, KeyRecaptchaURL)
	recaptchaSecret := config.GetStringValueWithNameSpace(hostname, KeyRecaptchaSecret)
	fmt.Println("url is: ", recaptchaURL, "secret is: ", recaptchaSecret)
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}
	request := url.Values{"secret": {recaptchaSecret}, "response": {in.Response}}
	resp, err := httpClient.PostForm(recaptchaURL, request)
	if err != nil {
		fmt.Println("http error is", err)
		return &npool.VerifyGoogleRecaptchaResponse{
			Info: false,
		}, xerrors.Errorf("fail to verify: %v", err)
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return &npool.VerifyGoogleRecaptchaResponse{
			Info: false,
		}, xerrors.Errorf("read response body error: %v", err)
	}

	return &npool.VerifyGoogleRecaptchaResponse{
		Info: true,
	}, nil
}
