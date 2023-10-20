package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterFulfiltaskQueryAPIResponse 核销单查询 API返回值
// alibaba.servicecenter.fulfiltask.query
//
// 当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
type AlibabaServicecenterFulfiltaskQueryAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterFulfiltaskQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterFulfiltaskQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterFulfiltaskQueryAPIResponseModel).Reset()
}

// AlibabaServicecenterFulfiltaskQueryAPIResponseModel is 核销单查询 成功返回结果
type AlibabaServicecenterFulfiltaskQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_fulfiltask_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *PagedResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterFulfiltaskQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaServicecenterFulfiltaskQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterFulfiltaskQueryAPIResponse)
	},
}

// GetAlibabaServicecenterFulfiltaskQueryAPIResponse 从 sync.Pool 获取 AlibabaServicecenterFulfiltaskQueryAPIResponse
func GetAlibabaServicecenterFulfiltaskQueryAPIResponse() *AlibabaServicecenterFulfiltaskQueryAPIResponse {
	return poolAlibabaServicecenterFulfiltaskQueryAPIResponse.Get().(*AlibabaServicecenterFulfiltaskQueryAPIResponse)
}

// ReleaseAlibabaServicecenterFulfiltaskQueryAPIResponse 将 AlibabaServicecenterFulfiltaskQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterFulfiltaskQueryAPIResponse(v *AlibabaServicecenterFulfiltaskQueryAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterFulfiltaskQueryAPIResponse.Put(v)
}
