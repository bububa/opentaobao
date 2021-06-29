package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取销退收货信息 APIResponse
taobao.wlb.wms.return.bill.get

通过订单号获取单个销退单收货信息。
*/
type TaobaoWlbWmsReturnBillGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsReturnBillGetResponse
}

type TaobaoWlbWmsReturnBillGetResponse struct {
    XMLName xml.Name `xml:"wlb_wms_return_bill_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 回退订单信息
    
    ReturnOrderInfo   *CainiaoReturnBillReturnorderinfo `json:"return_order_info,omitempty" xml:"return_order_info,omitempty"`

    
}
