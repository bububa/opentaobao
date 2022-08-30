package alihouse

import (
	"encoding/xml"

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

// AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponseModel is 提交小区专家 成功返回结果
type AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_community_broker_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
