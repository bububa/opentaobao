package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateDeleteAPIResponse 淘宝短信模板删除 API返回值
// taobao.jst.sms.template.delete
//
// 淘宝短信模板删除
type TaobaoJstSmsTemplateDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsTemplateDeleteAPIResponseModel
}

// TaobaoJstSmsTemplateDeleteAPIResponseModel is 淘宝短信模板删除 成功返回结果
type TaobaoJstSmsTemplateDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_template_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求成功
	RSuccess bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
	// 删除成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
