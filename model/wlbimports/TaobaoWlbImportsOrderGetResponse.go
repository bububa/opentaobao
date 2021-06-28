package wlbimports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流订单获取 APIResponse
taobao.wlb.imports.order.get

一般进口物流订单获取
*/
type TaobaoWlbImportsOrderGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_imports_order_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 物流订单信息
    
    Orders   []LocOrder `json:"orders,omitempty" xml:"