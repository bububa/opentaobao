package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointCollectionstandpointQueryAPIResponse 查询收藏口径 API返回值
// alibaba.legal.standpoint.collectionstandpoint.query
//
// 查询收藏口径
type AlibabaLegalStandpointCollectionstandpointQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointCollectionstandpointQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointCollectionstandpointQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointCollectionstandpointQueryAPIResponseModel).Reset()
}

// AlibabaLegalStandpointCollectionstandpointQueryAPIResponseModel is 查询收藏口径 成功返回结果
type AlibabaLegalStandpointCollectionstandpointQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_collectionstandpoint_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统错误
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 500
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 收藏口径
	Content *Page `json:"content,omitempty" xml:"content,omitempty"`
	// true
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointCollectionstandpointQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.Content = nil
	m.SuccessRes = false
}

var poolAlibabaLegalStandpointCollectionstandpointQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointCollectionstandpointQueryAPIResponse)
	},
}

// GetAlibabaLegalStandpointCollectionstandpointQueryAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointCollectionstandpointQueryAPIResponse
func GetAlibabaLegalStandpointCollectionstandpointQueryAPIResponse() *AlibabaLegalStandpointCollectionstandpointQueryAPIResponse {
	return poolAlibabaLegalStandpointCollectionstandpointQueryAPIResponse.Get().(*AlibabaLegalStandpointCollectionstandpointQueryAPIResponse)
}

// ReleaseAlibabaLegalStandpointCollectionstandpointQueryAPIResponse 将 AlibabaLegalStandpointCollectionstandpointQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointCollectionstandpointQueryAPIResponse(v *AlibabaLegalStandpointCollectionstandpointQueryAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointCollectionstandpointQueryAPIResponse.Put(v)
}
