package jwt

import (
	"encoding/base64"
	"encoding/json"
	"library-management/util"
	"log"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

type Payload struct {
	Iss string `json:"iss"`
	Exp string `json:"exp"`
	Iat string `json:"iat"`
	Sub string `json:"sub"`
	Typ string `json:"typ"`
	Uid string `json:"uid"`
}

// CreateHeader 创建Token头部
func CreateHeader() string {
	h := Header{
		Typ: "JWT",
		Alg: "HS256",
	}
	bytes, err := json.Marshal(h)
	if err != nil {
		log.Printf("marshal failed!序列化失败！错误信息：%v", err)
		return ""
	}
	header := base64.StdEncoding.EncodeToString(bytes)
	return header
}

// CreatePayload 创建Token载荷
func CreatePayload(username string, uid uint, t string) string {
	p := Payload{
		Iss: "HoshinoSyou",
		Exp: strconv.FormatInt(time.Now().Add(15*24*time.Hour).Unix(), 10),
		Iat: strconv.FormatInt(time.Now().Unix(), 10),
		Sub: username,
		Typ: t,
		Uid: strconv.FormatUint(uint64(uid), 10),
	}
	bytes, err := json.Marshal(p)
	if err != nil {
		log.Printf("marshal failed!序列化失败！错误信息：%v", err)
		return ""
	}
	payload := base64.StdEncoding.EncodeToString(bytes)
	return payload
}

// CreateSignature 创建签证
func CreateSignature(h, p string) string {
	str := strings.Join([]string{h, p}, ".")
	key := "HMACHoshinoSyouJWT"
	signature := util.Sha256(str, key)
	return signature
}

// Create 登陆后创建Token
func Create(username string, uid uint, t string) (token string) {
	header := CreateHeader()
	payload := CreatePayload(username, uid, t)
	signature := CreateSignature(header, payload)
	token = strings.Join([]string{header, payload}, ".") + "." + signature
	return token
}
