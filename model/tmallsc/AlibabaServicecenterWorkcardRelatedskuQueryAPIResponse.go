package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse 查询工单关联的服务项 API返回值
// alibaba.servicecenter.workcard.relatedsku.query
//
// 查询工单关联的服务项
type AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardRelatedskuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterWorkcardRelatedskuQueryAPIResponseModel).Reset()
}

// AlibabaServicecenterWorkcardRelatedskuQueryAPIResponseModel is 查询工单关联的服务项 成功返回结果
type AlibabaServicecenterWorkcardRelatedskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_relatedsku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaServicecenterWorkcardRelatedskuQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardRelatedskuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterWorkcardRelatedskuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse)
	},
}

// GetAlibabaServicecenterWorkcardRelatedskuQueryAPIResponse 从 sync.Pool 获取 AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse
func GetAlibabaServicecenterWorkcardRelatedskuQueryAPIResponse() *AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse {
	return poolAlibabaServicecenterWorkcardRelatedskuQueryAPIResponse.Get().(*AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse)
}

// ReleaseAlibabaServicecenterWorkcardRelatedskuQueryAPIResponse 将 AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterWorkcardRelatedskuQueryAPIResponse(v *AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterWorkcardRelatedskuQueryAPIResponse.Put(v)
}
