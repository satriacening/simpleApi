package helper

import (
	"encoding/base64"
)

func RandomStr(referal string) string {
	res := base64.StdEncoding.EncodeToString([]byte(referal))
	return (res[0:8])
}
