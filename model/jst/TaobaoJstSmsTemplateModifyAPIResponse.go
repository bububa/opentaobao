package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateModifyAPIResponse 淘宝短信模板修改 API返回值
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
type TaobaoJstSmsTemplateModifyAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsTemplateModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstSmsTemplateModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsTemplateModifyAPIResponseModel).Reset()
}

// TaobaoJstSmsTemplateModifyAPIResponseModel is 淘宝短信模板修改 成功返回结果
type TaobaoJstSmsTemplateModifyAPIResponseModel struct {
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

// Reset 清空结构体
func (m *TaobaoJstSmsTemplateModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RCode = ""
	m.Message = ""
	m.RSuccess = false
	m.Module = false
}

var poolTaobaoJstSmsTemplateModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsTemplateModifyAPIResponse)
	},
}

// GetTaobaoJstSmsTemplateModifyAPIResponse 从 sync.Pool 获取 TaobaoJstSmsTemplateModifyAPIResponse
func GetTaobaoJstSmsTemplateModifyAPIResponse() *TaobaoJstSmsTemplateModifyAPIResponse {
	return poolTaobaoJstSmsTemplateModifyAPIResponse.Get().(*TaobaoJstSmsTemplateModifyAPIResponse)
}

// ReleaseTaobaoJstSmsTemplateModifyAPIResponse 将 TaobaoJstSmsTemplateModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsTemplateModifyAPIResponse(v *TaobaoJstSmsTemplateModifyAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsTemplateModifyAPIResponse.Put(v)
}
