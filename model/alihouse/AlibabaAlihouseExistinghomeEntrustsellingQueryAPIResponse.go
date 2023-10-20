package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse 委托信息查询接口 API返回值
// alibaba.alihouse.existinghome.entrustselling.query
//
// 管家状态及房源信息接口
type AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponseModel is 委托信息查询接口 成功返回结果
type AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_entrustselling_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeEntrustsellingQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse
func GetAlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse() *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse {
	return poolAlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse.Get().(*AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse 将 AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse(v *AlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeEntrustsellingQueryAPIResponse.Put(v)
}
