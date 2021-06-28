package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单详情 APIResponse
alibaba.alihealth.nr.trade.order.get

阿里健康O2O，获取订单详情
*/
type AlibabaAlihealthNrTradeOrderGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alihealth_nr_trade_order_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 淘宝订单id
    
    OrderId   int64 `json:"order_id,omitempty" xml:"