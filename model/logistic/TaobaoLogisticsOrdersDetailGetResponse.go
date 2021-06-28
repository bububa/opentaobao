package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询物流订单,返回详细信息 APIResponse
taobao.logistics.orders.detail.get

查询物流订单的详细信息，涉及用户隐私字段。
*/
type TaobaoLogisticsOrdersDetailGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"logistics_orders_detail_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 获取的物流订单详情列表.返回的Shipping包含的具体信息为入参fields请求的字段信息.
    
    Shippings   []Shipping `json:"shippings,omitempty" xml:"