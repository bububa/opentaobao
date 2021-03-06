package tvpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayOrderPartnerpayAPIResponse tv支付第三方支付订单 API返回值
// taobao.tvpay.order.partnerpay
//
// tv支付第三方发起并支付订单（使用设备授权）
type TaobaoTvpayOrderPartnerpayAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayOrderPartnerpayAPIResponseModel
}

// TaobaoTvpayOrderPartnerpayAPIResponseModel is tv支付第三方支付订单 成功返回结果
type TaobaoTvpayOrderPartnerpayAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_order_partnerpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
