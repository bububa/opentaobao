package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse 提交小区专家 API返回值
// alibaba.alihouse.existinghome.community.broker.submit
//
// 提交小区专家
type AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponseModel is 提交小区专家 成功返回结果
type AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_community_broker_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse
func GetAlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse() *AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse {
	return poolAlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse.Get().(*AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse 将 AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse(v *AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse.Put(v)
}
