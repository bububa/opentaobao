package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmstemplatequeryAPIResponse 淘宝短信模板查询 API返回值
// taobao.jst.sms.template.query
//
// 淘宝短信模板查询
type TaobaojstsmstemplatequeryAPIResponse struct {
	model.CommonResponse
	TaobaojstsmstemplatequeryAPIResponseModel
}

// TaobaojstsmstemplatequeryAPIResponseModel is 淘宝短信模板查询 成功返回结果
type TaobaojstsmstemplatequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 错误消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Module *AccessBaseDto `json:"module,omitempty" xml:"module,omitempty"`
	// 请求成功
	RSuccess bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
}
