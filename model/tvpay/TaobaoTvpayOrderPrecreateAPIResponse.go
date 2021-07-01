package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付预下单 API返回值 
taobao.tvpay.order.precreate

tv支付预下单
*/
type TaobaoTvpayOrderPrecreateAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayOrderPrecreateAPIResponseModel
}

// tv支付预下单 成功返回结果
type TaobaoTvpayOrderPrecreateAPIResponseModel struct {
    XMLName xml.Name `xml:"tvpay_order_precreate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
