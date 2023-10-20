package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointCollectionAPIResponse 收藏|取消收藏 API返回值
// alibaba.legal.standpoint.standpoint.collection
//
// 收藏|取消收藏
type AlibabaLegalStandpointStandpointCollectionAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointStandpointCollectionAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointStandpointCollectionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointStandpointCollectionAPIResponseModel).Reset()
}

// AlibabaLegalStandpointStandpointCollectionAPIResponseModel is 收藏|取消收藏 成功返回结果
type AlibabaLegalStandpointStandpointCollectionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_standpoint_collection_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 500
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// true
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
	// true
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointStandpointCollectionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.SuccessRes = false
	m.Content = false
}

var poolAlibabaLegalStandpointStandpointCollectionAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointStandpointCollectionAPIResponse)
	},
}

// GetAlibabaLegalStandpointStandpointCollectionAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointStandpointCollectionAPIResponse
func GetAlibabaLegalStandpointStandpointCollectionAPIResponse() *AlibabaLegalStandpointStandpointCollectionAPIResponse {
	return poolAlibabaLegalStandpointStandpointCollectionAPIResponse.Get().(*AlibabaLegalStandpointStandpointCollectionAPIResponse)
}

// ReleaseAlibabaLegalStandpointStandpointCollectionAPIResponse 将 AlibabaLegalStandpointStandpointCollectionAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointStandpointCollectionAPIResponse(v *AlibabaLegalStandpointStandpointCollectionAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointStandpointCollectionAPIResponse.Put(v)
}
