package jst

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoJstSmsSignnameModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsSignnameModifyAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *TaobaoJstSmsSignnameModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RCode = ""
	m.Message = ""
	m.RSuccess = false
	m.Module = false
}

var poolTaobaoJstSmsSignnameModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsSignnameModifyAPIResponse)
	},
}

// GetTaobaoJstSmsSignnameModifyAPIResponse 从 sync.Pool 获取 TaobaoJstSmsSignnameModifyAPIResponse
func GetTaobaoJstSmsSignnameModifyAPIResponse() *TaobaoJstSmsSignnameModifyAPIResponse {
	return poolTaobaoJstSmsSignnameModifyAPIResponse.Get().(*TaobaoJstSmsSignnameModifyAPIResponse)
}

// ReleaseTaobaoJstSmsSignnameModifyAPIResponse 将 TaobaoJstSmsSignnameModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsSignnameModifyAPIResponse(v *TaobaoJstSmsSignnameModifyAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsSignnameModifyAPIResponse.Put(v)
}
