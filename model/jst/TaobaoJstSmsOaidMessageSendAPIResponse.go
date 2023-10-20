package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmsoaidmessagesendAPIResponse 基于OAID的短信发送接口 API返回值
// taobao.jst.sms.oaid.message.send
//
// 基于OAID的短信发送接口
type TaobaojstsmsoaidmessagesendAPIResponse struct {
	model.CommonResponse
	TaobaojstsmsoaidmessagesendAPIResponseModel
}

// TaobaojstsmsoaidmessagesendAPIResponseModel is 基于OAID的短信发送接口 成功返回结果
type TaobaojstsmsoaidmessagesendAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_oaid_message_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 短信拓展码
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// top请求id
	ReqId string `json:"req_id,omitempty" xml:"req_id,omitempty"`
}
