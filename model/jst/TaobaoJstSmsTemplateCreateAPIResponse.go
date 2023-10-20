package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateCreateAPIResponse 淘宝短信模板创建 API返回值
// taobao.jst.sms.template.create
//
// 聚石塔短信模板创建
type TaobaoJstSmsTemplateCreateAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsTemplateCreateAPIResponseModel
}

// TaobaoJstSmsTemplateCreateAPIResponseModel is 淘宝短信模板创建 成功返回结果
type TaobaoJstSmsTemplateCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_template_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 模板CODE
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 失败原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求成功
	RSuccess bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
}
