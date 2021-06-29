package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取销退收货信息 API返回值 
taobao.wlb.wms.return.bill.get

通过订单号获取单个销退单收货信息。
*/
type TaobaoWlbWmsReturnBillGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsReturnBillGetResponse
}

// 获取销退收货信息 成功返回结果
type TaobaoWlbWmsReturnBillGetResponse struct {
    XMLName xml.Name `xml:"wlb_wms_return_bill_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 回退订单信息
    ReturnOrderInfo   *CainiaoReturnBillReturnorderinfo `json:"return_order_info,omitempty" xml:"return_order_info,omitempty"`
}
