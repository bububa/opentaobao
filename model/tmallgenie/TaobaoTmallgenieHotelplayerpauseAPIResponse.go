package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmallgeniehotelplayerpauseAPIResponse 天猫精灵酒店播放暂停 API返回值
// taobao.tmallgenie.hotelplayerpause
//
// 酒店推送指令给天猫精灵停止播放音乐
type TaobaotmallgeniehotelplayerpauseAPIResponse struct {
	model.CommonResponse
	TaobaotmallgeniehotelplayerpauseAPIResponseModel
}

// TaobaotmallgeniehotelplayerpauseAPIResponseModel is 天猫精灵酒店播放暂停 成功返回结果
type TaobaotmallgeniehotelplayerpauseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmallgenie_hotelplayerpause_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// statusCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
