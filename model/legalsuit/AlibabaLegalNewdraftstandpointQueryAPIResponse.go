package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalNewdraftstandpointQueryAPIResponse 未采纳口径查询(新) API返回值
// alibaba.legal.newdraftstandpoint.query
//
// 未采纳口径查询(新)
type AlibabaLegalNewdraftstandpointQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLegalNewdraftstandpointQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalNewdraftstandpointQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalNewdraftstandpointQueryAPIResponseModel).Reset()
}

// AlibabaLegalNewdraftstandpointQueryAPIResponseModel is 未采纳口径查询(新) 成功返回结果
type AlibabaLegalNewdraftstandpointQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_newdraftstandpoint_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 100
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回状态
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 分页对象
	Content *Page `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalNewdraftstandpointQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.Content = nil
	m.SuccessRes = false
}

var poolAlibabaLegalNewdraftstandpointQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalNewdraftstandpointQueryAPIResponse)
	},
}

// GetAlibabaLegalNewdraftstandpointQueryAPIResponse 从 sync.Pool 获取 AlibabaLegalNewdraftstandpointQueryAPIResponse
func GetAlibabaLegalNewdraftstandpointQueryAPIResponse() *AlibabaLegalNewdraftstandpointQueryAPIResponse {
	return poolAlibabaLegalNewdraftstandpointQueryAPIResponse.Get().(*AlibabaLegalNewdraftstandpointQueryAPIResponse)
}

// ReleaseAlibabaLegalNewdraftstandpointQueryAPIResponse 将 AlibabaLegalNewdraftstandpointQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalNewdraftstandpointQueryAPIResponse(v *AlibabaLegalNewdraftstandpointQueryAPIResponse) {
	v.Reset()
	poolAlibabaLegalNewdraftstandpointQueryAPIResponse.Put(v)
}
