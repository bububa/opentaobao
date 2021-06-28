package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单id获取单条订单详情 APIResponse
alibaba.alihealth.nr.trade.order.getorderdetail

阿里健康O2O，获取订单详情，修复组合商品价格精度问题
*/
type AlibabaAlihealthNrTradeOrderGetorderdetailAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alihealth_nr_trade_order_getorderdetail_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 高德，本地商户
    
    OutStoreId   string `json:"out_store_id,omitempty" xml:"