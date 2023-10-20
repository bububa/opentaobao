package tmallcar

import (
	"encoding/xml"

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

// TmallAliautoFulfillmentContractSignAPIResponseModel is 合同签署 成功返回结果
type TmallAliautoFulfillmentContractSignAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_fulfillment_contract_sign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}
