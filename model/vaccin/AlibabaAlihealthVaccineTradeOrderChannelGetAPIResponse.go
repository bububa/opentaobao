package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineTradeOrderChannelGetAPIResponse 通过订单ID与卖家ID获取订单渠道 API返回值
// alibaba.alihealth.vaccine.trade.order.channel.get
//
// 通过订单ID与卖家ID获取订单渠道
type AlibabaAlihealthVaccineTradeOrderChannelGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthVaccineTradeOrderChannelGetAPIResponseModel
}

// AlibabaAlihealthVaccineTradeOrderChannelGetAPIResponseModel is 通过订单ID与卖家ID获取订单渠道 成功返回结果
type AlibabaAlihealthVaccineTradeOrderChannelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_vaccine_trade_order_channel_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务响应code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 订单渠道
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 业务错误信息描述
	BizMessage string `json:"biz_message,omitempty" xml:"biz_message,omitempty"`
	// 业务成功状态
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
