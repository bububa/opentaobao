package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaicontentbusinesssupplychargeAPIResponse 供销商品充值接口 API返回值
// alibaba.ai.content.business.supply.charge
//
// 供销商品充值接口
type AlibabaaicontentbusinesssupplychargeAPIResponse struct {
	model.CommonResponse
	AlibabaaicontentbusinesssupplychargeAPIResponseModel
}

// AlibabaaicontentbusinesssupplychargeAPIResponseModel is 供销商品充值接口 成功返回结果
type AlibabaaicontentbusinesssupplychargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ai_content_business_supply_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaaicontentbusinesssupplychargeBizResult `json:"result,omitempty" xml:"result,omitempty"`
}
