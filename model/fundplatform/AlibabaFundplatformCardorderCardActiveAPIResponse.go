package fundplatform

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardorderCardActiveAPIResponse 储值卡激活 API返回值
// alibaba.fundplatform.cardorder.card.active
//
// 储值卡接货接口，可以通过外部订单号或者卡号进行批量激活。如果储值卡已经被激活过仍然幂等返回成功。资金平台保证批量激活时一定全部成功或全部失败。
// 如果批量激活储值卡时，如果部分储值卡处于已激活，部分储值卡处于未激活，则会返回失败
type AlibabaFundplatformCardorderCardActiveAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformCardorderCardActiveAPIResponseModel
}

// AlibabaFundplatformCardorderCardActiveAPIResponseModel is 储值卡激活 成功返回结果
type AlibabaFundplatformCardorderCardActiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorder_card_active_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *CardActiveResponse `json:"result,omitempty" xml:"result,omitempty"`
}
