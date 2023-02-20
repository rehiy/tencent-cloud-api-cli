package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func AuthCode(service, payload, secretId, secretKey string, timestamp int64) string {

	algorithm := "TC3-HMAC-SHA256"
	host := service + ".tencentcloudapi.com"

	// step 1: build canonical request string

	httpRequestMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	canonicalHeaders := "content-type:application/json; charset=utf-8\n" + "host:" + host + "\n"
	signedHeaders := "content-type;host"
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload,
	)

	// step 2: build string to sign

	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf(
		"%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest,
	)

	// step 3: sign string

	secretDate := hmacsha256(date, "TC3"+secretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))

	// step 4: build authorization

	return fmt.Sprintf(
		"%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secretId,
		credentialScope,
		signedHeaders,
		signature,
	)

}

func sha256hex(s string) string {

	b := sha256.Sum256([]byte(s))

	return hex.EncodeToString(b[:])

}

func hmacsha256(s, key string) string {

	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))

	return string(hashed.Sum(nil))

}
