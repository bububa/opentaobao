package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询物流宝订单 API返回值 
taobao.wlb.order.page.get

分页查询物流宝订单
*/
type TaobaoWlbOrderPageGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderPageGetAPIResponseModel
}

// 分页查询物流宝订单 成功返回结果
type TaobaoWlbOrderPageGetAPIResponseModel struct {
    XMLName xml.Name `xml:"wlb_order_page_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 物流宝订单对象
    OrderList   []WlbOrder `json:"order_list,omitempty" xml:"order_list>wlb_order,omitempty"`
    // 总条数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
