package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoFulfillmentContractSignAPIResponse 合同签署 API返回值
// tmall.aliauto.fulfillment.contract.sign
//
// 商家回传用户签署的合同信息
type TmallAliautoFulfillmentContractSignAPIResponse struct {
	model.CommonResponse
	TmallAliautoFulfillmentContractSignAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoFulfillmentContractSignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoFulfillmentContractSignAPIResponseModel).Reset()
}

// TmallAliautoFulfillmentContractSignAPIResponseModel is 合同签署 成功返回结果
type TmallAliautoFulfillmentContractSignAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_fulfillment_contract_sign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoFulfillmentContractSignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoFulfillmentContractSignAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoFulfillmentContractSignAPIResponse)
	},
}

// GetTmallAliautoFulfillmentContractSignAPIResponse 从 sync.Pool 获取 TmallAliautoFulfillmentContractSignAPIResponse
func GetTmallAliautoFulfillmentContractSignAPIResponse() *TmallAliautoFulfillmentContractSignAPIResponse {
	return poolTmallAliautoFulfillmentContractSignAPIResponse.Get().(*TmallAliautoFulfillmentContractSignAPIResponse)
}

// ReleaseTmallAliautoFulfillmentContractSignAPIResponse 将 TmallAliautoFulfillmentContractSignAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoFulfillmentContractSignAPIResponse(v *TmallAliautoFulfillmentContractSignAPIResponse) {
	v.Reset()
	poolTmallAliautoFulfillmentContractSignAPIResponse.Put(v)
}
