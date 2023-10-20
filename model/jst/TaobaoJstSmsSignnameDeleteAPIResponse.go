package jst

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoJstSmsSignnameDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsSignnameDeleteAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *TaobaoJstSmsSignnameDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RCode = ""
	m.Message = ""
	m.RSuccess = false
	m.Module = false
}

var poolTaobaoJstSmsSignnameDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsSignnameDeleteAPIResponse)
	},
}

// GetTaobaoJstSmsSignnameDeleteAPIResponse 从 sync.Pool 获取 TaobaoJstSmsSignnameDeleteAPIResponse
func GetTaobaoJstSmsSignnameDeleteAPIResponse() *TaobaoJstSmsSignnameDeleteAPIResponse {
	return poolTaobaoJstSmsSignnameDeleteAPIResponse.Get().(*TaobaoJstSmsSignnameDeleteAPIResponse)
}

// ReleaseTaobaoJstSmsSignnameDeleteAPIResponse 将 TaobaoJstSmsSignnameDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsSignnameDeleteAPIResponse(v *TaobaoJstSmsSignnameDeleteAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsSignnameDeleteAPIResponse.Put(v)
}
