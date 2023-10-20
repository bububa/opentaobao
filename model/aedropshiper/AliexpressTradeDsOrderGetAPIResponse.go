package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresstradedsordergetAPIResponse 买家查询订单详情 API返回值
// aliexpress.trade.ds.order.get
//
// 买家查询订单详情，用于dropshipper
type AliexpresstradedsordergetAPIResponse struct {
	model.CommonResponse
	AliexpresstradedsordergetAPIResponseModel
}

// AliexpresstradedsordergetAPIResponseModel is 买家查询订单详情 成功返回结果
type AliexpresstradedsordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_trade_ds_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单信息
	Result *AeopOrderInfo `json:"result,omitempty" xml:"result,omitempty"`
}
