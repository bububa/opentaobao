package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmsmessageshorturlcreateAPIResponse 聚石塔营销效果短链生成 API返回值
// taobao.jst.sms.message.shorturl.create
//
// 聚石塔生成淘短链接口
type TaobaojstsmsmessageshorturlcreateAPIResponse struct {
	model.CommonResponse
	TaobaojstsmsmessageshorturlcreateAPIResponseModel
}

// TaobaojstsmsmessageshorturlcreateAPIResponseModel is 聚石塔营销效果短链生成 成功返回结果
type TaobaojstsmsmessageshorturlcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_message_shorturl_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 请求ID
	RRequestId string `json:"r_request_id,omitempty" xml:"r_request_id,omitempty"`
	// 短链值(短链的有效期为2个月)
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 请求失败时的错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功
	RSuccess bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
}
