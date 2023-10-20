package kclub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKclubKcQueryknowledgeAPIResponse 知识云-通用知识查询服务 API返回值
// alibaba.kclub.kc.queryknowledge
//
// 知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
type AlibabaKclubKcQueryknowledgeAPIResponse struct {
	model.CommonResponse
	AlibabaKclubKcQueryknowledgeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaKclubKcQueryknowledgeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaKclubKcQueryknowledgeAPIResponseModel).Reset()
}

// AlibabaKclubKcQueryknowledgeAPIResponseModel is 知识云-通用知识查询服务 成功返回结果
type AlibabaKclubKcQueryknowledgeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kclub_kc_queryknowledge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaKclubKcQueryknowledgeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaKclubKcQueryknowledgeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaKclubKcQueryknowledgeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaKclubKcQueryknowledgeAPIResponse)
	},
}

// GetAlibabaKclubKcQueryknowledgeAPIResponse 从 sync.Pool 获取 AlibabaKclubKcQueryknowledgeAPIResponse
func GetAlibabaKclubKcQueryknowledgeAPIResponse() *AlibabaKclubKcQueryknowledgeAPIResponse {
	return poolAlibabaKclubKcQueryknowledgeAPIResponse.Get().(*AlibabaKclubKcQueryknowledgeAPIResponse)
}

// ReleaseAlibabaKclubKcQueryknowledgeAPIResponse 将 AlibabaKclubKcQueryknowledgeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaKclubKcQueryknowledgeAPIResponse(v *AlibabaKclubKcQueryknowledgeAPIResponse) {
	v.Reset()
	poolAlibabaKclubKcQueryknowledgeAPIResponse.Put(v)
}
