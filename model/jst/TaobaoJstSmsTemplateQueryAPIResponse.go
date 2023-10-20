package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateQueryAPIResponse 淘宝短信模板查询 API返回值
// taobao.jst.sms.template.query
//
// 淘宝短信模板查询
type TaobaoJstSmsTemplateQueryAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsTemplateQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstSmsTemplateQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsTemplateQueryAPIResponseModel).Reset()
}

// TaobaoJstSmsTemplateQueryAPIResponseModel is 淘宝短信模板查询 成功返回结果
type TaobaoJstSmsTemplateQueryAPIResponseModel struct {
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

// Reset 清空结构体
func (m *TaobaoJstSmsTemplateQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RCode = ""
	m.Message = ""
	m.Module = nil
	m.RSuccess = false
}

var poolTaobaoJstSmsTemplateQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsTemplateQueryAPIResponse)
	},
}

// GetTaobaoJstSmsTemplateQueryAPIResponse 从 sync.Pool 获取 TaobaoJstSmsTemplateQueryAPIResponse
func GetTaobaoJstSmsTemplateQueryAPIResponse() *TaobaoJstSmsTemplateQueryAPIResponse {
	return poolTaobaoJstSmsTemplateQueryAPIResponse.Get().(*TaobaoJstSmsTemplateQueryAPIResponse)
}

// ReleaseTaobaoJstSmsTemplateQueryAPIResponse 将 TaobaoJstSmsTemplateQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsTemplateQueryAPIResponse(v *TaobaoJstSmsTemplateQueryAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsTemplateQueryAPIResponse.Put(v)
}
