package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmemoalarmcreateAPIResponse 天猫精灵闹钟创建 API返回值
// taobao.ailab.aicloud.top.memo.alarm.create
//
// 天猫精灵闹钟创建
type TaobaoailabaicloudtopmemoalarmcreateAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopmemoalarmcreateAPIResponseModel
}

// TaobaoailabaicloudtopmemoalarmcreateAPIResponseModel is 天猫精灵闹钟创建 成功返回结果
type TaobaoailabaicloudtopmemoalarmcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_memo_alarm_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 闹钟ID
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 状态码
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
