package missevan

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/google/uuid"
)

func getBaseCookies() (cookies [2]*http.Cookie) {
	resp, _ := http.Get(UserInfoUrl)
	cks := resp.Cookies()
	for _, ck := range cks {
		switch ck.Name {
		case "FM_SESS":
			cookies[FM_SESS] = ck
		case "FM_SESS.sig":
			cookies[FM_SESS_SIG] = ck
		}
	}
	return
}

const secret = "f8Yw1M3u0e5MV7z34PcbY9wgP56YwJ"

func buildAccessTokenRequest(phone int, password string) *http.Request {
	payload := fmt.Sprintf("account=%d&password=%s&region=CN", phone, password)
	xmdate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	xmnonce := "247b7772-63e4-4185-b774-17d5ecf821e6"
	equipID := "52e76e9a-7592-43af-ae72-8ec755094e41"
	baseCookies := getBaseCookies()
	req, _ := http.NewRequest(http.MethodPost, LoginUrl, strings.NewReader(payload))

	req.Header.Add("Host", "app.missevan.com")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-Hans-CN;q=1, en-CN;q=0.9, ja-CN;q=0.8")
	req.Header.Add("User-Agent", "MissEvanApp/4.7.4 (iOS;15.1;iPhone11,6)")
	req.Header.Add("X-M-Date", xmdate)
	req.Header.Add("X-M-Nonce", xmnonce)
	req.AddCookie(baseCookies[FM_SESS])
	req.AddCookie(baseCookies[FM_SESS_SIG])
	req.Header.Add("Cookie", "equip_id="+equipID)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", encodeAuth(xmdate, xmnonce, equipID, payload))
	return req
}

func encodeAuth(xmdate, xmnonce, equipID string, payload string) string {
	sum := sha256.Sum256([]byte(payload))
	enc := base64.StdEncoding.EncodeToString(sum[:])
	_url := strings.ReplaceAll(LoginUrl, ":", "%3A")
	str := http.MethodPost + "\n" +
		_url + "\n" +
		"" + "\n" +
		"equip_id:" + equipID + "\n" +
		"x-m-date:" + xmdate + "\n" +
		"x-m-nonce:" + xmnonce + "\n" +
		enc + "\n"
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	dst := h.Sum(nil)
	enc = base64.StdEncoding.EncodeToString((dst))
	return "MissEvan " + enc
}

func BrotliDecompress(data []byte) ([]byte, error) {
	r := brotli.NewReader(bytes.NewReader(data[4:]))
	b, _ := ioutil.ReadAll(r)
	return b, nil
}

func MessageID() string {
	u := uuid.NewString()
	return "3" + u[1:]
}

func SafeMessage(msg string) string {
	return strings.Join(strings.Split(msg, ""), "\u200B")
}

func NewToken(phone int, password string) (string, error) {
	fmResp := FMResp{}
	c := &http.Client{}
	req := buildAccessTokenRequest(phone, password)
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(bs, &fmResp); err != nil {
		return "", err
	}
	if fmResp.Code != 0 || fmResp.Success != true {
		return "", errors.New(string(bs))
	}
	return fmResp.Info.Token, nil
}
