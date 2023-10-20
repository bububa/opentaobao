package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmstemplatemodifyAPIResponse 淘宝短信模板修改 API返回值
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
type TaobaojstsmstemplatemodifyAPIResponse struct {
	model.CommonResponse
	TaobaojstsmstemplatemodifyAPIResponseModel
}

// TaobaojstsmstemplatemodifyAPIResponseModel is 淘宝短信模板修改 成功返回结果
type TaobaojstsmstemplatemodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_template_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求成功
	RSuccess bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
	// 修改成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
