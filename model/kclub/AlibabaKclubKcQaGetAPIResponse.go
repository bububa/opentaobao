package kclub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKclubKcQaGetAPIResponse 知识云-查询单个知识详情 API返回值
// alibaba.kclub.kc.qa.get
//
// 知识云-查询单个知识详情。通过租户id、问题id查询问题详情
type AlibabaKclubKcQaGetAPIResponse struct {
	model.CommonResponse
	AlibabaKclubKcQaGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaKclubKcQaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaKclubKcQaGetAPIResponseModel).Reset()
}

// AlibabaKclubKcQaGetAPIResponseModel is 知识云-查询单个知识详情 成功返回结果
type AlibabaKclubKcQaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kclub_kc_qa_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaKclubKcQaGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaKclubKcQaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaKclubKcQaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcQaGetAPIResponse)
	},
}

// GetAlibabaKclubKcQaGetAPIResponse 从 sync.Pool 获取 AlibabaKclubKcQaGetAPIResponse
func GetAlibabaKclubKcQaGetAPIResponse() *AlibabaKclubKcQaGetAPIResponse {
	return poolAlibabaKclubKcQaGetAPIResponse.Get().(*AlibabaKclubKcQaGetAPIResponse)
}

// ReleaseAlibabaKclubKcQaGetAPIResponse 将 AlibabaKclubKcQaGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaKclubKcQaGetAPIResponse(v *AlibabaKclubKcQaGetAPIResponse) {
	v.Reset()
	poolAlibabaKclubKcQaGetAPIResponse.Put(v)
}
