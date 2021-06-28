package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询物流宝订单 APIResponse
taobao.wlb.order.page.get

分页查询物流宝订单
*/
type TaobaoWlbOrderPageGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_order_page_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 物流宝订单对象
    
    OrderList   []WlbOrder `json:"order_list,omitempty" xml:"