package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameModifyAPIResponse 淘宝短信签名修改 API返回值
// taobao.jst.sms.signname.modify
//
// 淘宝短信签名修改，只能修改还未被审核的签名。
type TaobaoJstSmsSignnameModifyAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsSignnameModifyAPIResponseModel
}

// TaobaoJstSmsSignnameModifyAPIResponseModel is 淘宝短信签名修改 成功返回结果
type TaobaoJstSmsSignnameModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_signname_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误CODE
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 失败原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求是否成功
	RSuccess bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
	// 请求成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
