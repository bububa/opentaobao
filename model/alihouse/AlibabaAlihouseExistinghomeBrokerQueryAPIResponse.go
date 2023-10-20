package alihouse

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrokerQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeBrokerQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeBrokerQueryAPIResponseModel is 根据外部经纪人ID查询 成功返回结果
type AlibabaAlihouseExistinghomeBrokerQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_broker_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeBrokerQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrokerQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeBrokerQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrokerQueryAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeBrokerQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrokerQueryAPIResponse
func GetAlibabaAlihouseExistinghomeBrokerQueryAPIResponse() *AlibabaAlihouseExistinghomeBrokerQueryAPIResponse {
	return poolAlibabaAlihouseExistinghomeBrokerQueryAPIResponse.Get().(*AlibabaAlihouseExistinghomeBrokerQueryAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerQueryAPIResponse 将 AlibabaAlihouseExistinghomeBrokerQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrokerQueryAPIResponse(v *AlibabaAlihouseExistinghomeBrokerQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrokerQueryAPIResponse.Put(v)
}
