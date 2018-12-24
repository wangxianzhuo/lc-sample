package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	"github.com/mozillazg/request"
)

const (
	systemVer = "1.0"
	appSecret = ""
	appID     = ""
)

var deviceID = ""
var adminToken = ""

func main() {
	accessToken()
	// deviceList()
	// bindDeviceInfo(deviceID)
	// checkDeviceBindOrNot()
	// setAllStorageStrategy(deviceID, "on")
	// getStorageStrategy(deviceID)
	// queryCloudRecordCallNum(35)
	// setStorageStrategy(deviceID, "on")
	// openCloudRecord(deviceID, 35)
	// queryCloudRecordBitmap(deviceID, 2018, 12)
	// queryLocalRecordNum(deviceID, "2018-12-01 00:00:00", "2018-12-31 00:00:00")
	// queryLocalRecords(deviceID, "2018-12-01 00:00:00", "2018-12-31 00:00:00", "1-30")
	// queryCloudRecordNum(deviceID, "2018-12-01 00:00:00", "2018-12-31 00:00:00")
	// queryCloudRecords(deviceID, "2018-12-01 00:00:00", "2018-12-31 00:00:00", "1-30")
}

func queryLocalRecords(deviceID, beginTime, endTime, queryRange string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/queryLocalRecords", map[string]interface{}{
		"token":      adminToken,
		"deviceId":   deviceID,
		"channelId":  "0",
		"beginTime":  beginTime,
		"endTime":    endTime,
		"queryRange": queryRange,
	})
}

func queryLocalRecordNum(deviceID, beginTime, endTime string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/queryLocalRecordNum", map[string]interface{}{
		"token":     adminToken,
		"deviceId":  deviceID,
		"channelId": "0",
		"beginTime": beginTime,
		"endTime":   endTime,
	})
}

func queryLocalRecordBitmap(deviceID string, year, month int) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/queryLocalRecordBitmap", map[string]interface{}{
		"token":     adminToken,
		"deviceId":  deviceID,
		"channelId": "0",
		"year":      year,
		"month":     month,
	})
}

func queryCloudRecords(deviceID, beginTime, endTime, queryRange string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/queryCloudRecords", map[string]interface{}{
		"token":      adminToken,
		"deviceId":   deviceID,
		"channelId":  "0",
		"beginTime":  beginTime,
		"endTime":    endTime,
		"queryRange": queryRange,
	})
}

func queryCloudRecordNum(deviceID, beginTime, endTime string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/queryCloudRecordNum", map[string]interface{}{
		"token":     adminToken,
		"deviceId":  deviceID,
		"channelId": "0",
		"beginTime": beginTime,
		"endTime":   endTime,
	})
}

func queryCloudRecordBitmap(deviceID string, year, month int) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/queryCloudRecordBitmap", map[string]interface{}{
		"token":     adminToken,
		"deviceId":  deviceID,
		"channelId": "0",
		"year":      year,
		"month":     month,
	})
}

func deviceList() error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/deviceList", map[string]interface{}{
		"token":      adminToken,
		"queryRange": "1-10",
	})
}
func setAllStorageStrategy(deviceID string, status string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/setAllStorageStrategy", map[string]interface{}{
		"deviceId":  deviceID,
		"token":     adminToken,
		"channelId": "0",
		"status":    status,
	})
}
func openCloudRecord(deviceID string, strategyID int) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/openCloudRecord", map[string]interface{}{
		"deviceId":   deviceID,
		"token":      adminToken,
		"channelId":  "0",
		"strategyId": strategyID,
	})
}
func setStorageStrategy(deviceID string, status string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/setStorageStrategy", map[string]interface{}{
		"deviceId":  deviceID,
		"token":     adminToken,
		"channelId": "0",
		"status":    status,
	})
}
func getStorageStrategy(deviceID string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/getStorageStrategy", map[string]interface{}{
		"deviceId":  deviceID,
		"token":     adminToken,
		"channelId": "0",
	})
}

func queryCloudRecordCallNum(strategyID int) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/queryCloudRecordCallNum", map[string]interface{}{
		"token":      adminToken,
		"strategyId": strategyID,
	})
}

func accessToken() error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/accessToken", map[string]interface{}{})
}

func checkDeviceBindOrNot(deviceID string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/checkDeviceBindOrNot", map[string]interface{}{
		"deviceId": deviceID,
		"token":    adminToken,
	})
}

func bindDeviceInfo(deviceID string) error {
	return lcRequest("https://openapi.lechange.cn:443/openapi/bindDeviceInfo", map[string]interface{}{
		"deviceId": deviceID,
		"token":    adminToken,
	})
}

func lcRequest(url string, params map[string]interface{}) error {
	c := new(http.Client)
	req := request.NewRequest(c)
	data, _ := json.Marshal(newReqData(params))
	req.Body = bytes.NewReader(data)
	resp, err := req.Post(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 打印response body
	d, _ := resp.Content()
	var out bytes.Buffer
	json.Indent(&out, []byte(d), "", "  ")
	fmt.Println(out.String())
	return nil
}

type reqSystem struct {
	Ver   string `json:"ver"`
	Sign  string `json:"sign"`
	AppID string `json:"appId"`
	Time  string `json:"time"`
	Nonce string `json:"nonce"`
}

func newReqSystem() reqSystem {
	nonce := generateNonce()
	time := fmt.Sprint(time.Now().Unix())
	sign := generateSign(time, nonce, appSecret)

	return reqSystem{
		Ver:   systemVer,
		Sign:  sign,
		AppID: appID,
		Time:  time,
		Nonce: nonce,
	}
}

type respSystem struct {
	Msg  string                 `json:"msg"`
	Code string                 `json:"code"`
	Data map[string]interface{} `json:"data"`
}

type reqData struct {
	ID     string                 `json:"id"`
	System reqSystem              `json:"system"`
	Params map[string]interface{} `json:"params"`
}

func newReqData(params map[string]interface{}) reqData {
	rand.Seed(time.Now().Unix())
	id := rand.Intn(100)
	return reqData{
		ID:     fmt.Sprint(id),
		System: newReqSystem(),
		Params: params,
	}
}

type respData struct {
	ID     string     `json:"id"`
	System respSystem `json:"system"`
}

func generateNonce() string {
	u := uuid.NewV4()
	buf := make([]byte, 32)

	hex.Encode(buf[0:8], u[0:4])
	hex.Encode(buf[8:12], u[4:6])
	hex.Encode(buf[12:16], u[6:8])
	hex.Encode(buf[16:20], u[8:10])
	hex.Encode(buf[20:], u[10:])

	return string(buf)
}

func generateSign(time, nonce, appSecret string) string {
	raw := fmt.Sprintf("time:%s,nonce:%s,appSecret:%s", time, nonce, appSecret)
	return fmt.Sprintf("%x", md5.Sum([]byte(raw)))
}
