package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeCustomizationRefundEnableAPIResponse 定制订单设置允许仅退款 API返回值
// taobao.opentrade.customization.refund.enable
//
// 定制订单设置允许仅退款
type TaobaoOpentradeCustomizationRefundEnableAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeCustomizationRefundEnableAPIResponseModel
}

// TaobaoOpentradeCustomizationRefundEnableAPIResponseModel is 定制订单设置允许仅退款 成功返回结果
type TaobaoOpentradeCustomizationRefundEnableAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_customization_refund_enable_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否设置成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
