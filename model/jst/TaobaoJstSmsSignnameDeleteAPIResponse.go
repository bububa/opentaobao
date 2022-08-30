package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameDeleteAPIResponse 淘宝短信签名删除 API返回值
// taobao.jst.sms.signname.delete
//
// 淘宝短信签名删除
type TaobaoJstSmsSignnameDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsSignnameDeleteAPIResponseModel
}

// TaobaoJstSmsSignnameDeleteAPIResponseModel is 淘宝短信签名删除 成功返回结果
type TaobaoJstSmsSignnameDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_signname_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 错误原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求成功
	RSuccess bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
	// 删除成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
