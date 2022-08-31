package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerQueryAPIResponse 根据外部经纪人ID查询 API返回值
// alibaba.alihouse.existinghome.broker.query
//
// 根据外部经纪人ID查询
type AlibabaAlihouseExistinghomeBrokerQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeBrokerQueryAPIResponseModel
}

// AlibabaAlihouseExistinghomeBrokerQueryAPIResponseModel is 根据外部经纪人ID查询 成功返回结果
type AlibabaAlihouseExistinghomeBrokerQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_broker_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeBrokerQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
