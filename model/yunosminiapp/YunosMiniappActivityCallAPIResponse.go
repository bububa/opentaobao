package yunosminiapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosminiappactivitycallAPIResponse 调用活动接口 API返回值
// yunos.miniapp.activity.call
//
// 用于小程序调用活动接口
type YunosminiappactivitycallAPIResponse struct {
	model.CommonResponse
	YunosminiappactivitycallAPIResponseModel
}

// YunosminiappactivitycallAPIResponseModel is 调用活动接口 成功返回结果
type YunosminiappactivitycallAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_miniapp_activity_call_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *YunosminiappactivitycallResult `json:"result,omitempty" xml:"result,omitempty"`
}
