package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameCreateAPIResponse 淘宝短信签名创建接口 API返回值
// taobao.jst.sms.signname.create
//
// 聚石塔短信签名创建接口
type TaobaoJstSmsSignnameCreateAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsSignnameCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstSmsSignnameCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsSignnameCreateAPIResponseModel).Reset()
}

// TaobaoJstSmsSignnameCreateAPIResponseModel is 淘宝短信签名创建接口 成功返回结果
type TaobaoJstSmsSignnameCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_signname_create_response"`
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

// Reset 清空结构体
func (m *TaobaoJstSmsSignnameCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RCode = ""
	m.Message = ""
	m.RSuccess = false
	m.Module = false
}

var poolTaobaoJstSmsSignnameCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsSignnameCreateAPIResponse)
	},
}

// GetTaobaoJstSmsSignnameCreateAPIResponse 从 sync.Pool 获取 TaobaoJstSmsSignnameCreateAPIResponse
func GetTaobaoJstSmsSignnameCreateAPIResponse() *TaobaoJstSmsSignnameCreateAPIResponse {
	return poolTaobaoJstSmsSignnameCreateAPIResponse.Get().(*TaobaoJstSmsSignnameCreateAPIResponse)
}

// ReleaseTaobaoJstSmsSignnameCreateAPIResponse 将 TaobaoJstSmsSignnameCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsSignnameCreateAPIResponse(v *TaobaoJstSmsSignnameCreateAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsSignnameCreateAPIResponse.Put(v)
}
