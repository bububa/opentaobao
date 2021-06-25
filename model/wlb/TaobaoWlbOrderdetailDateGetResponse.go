package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按照日期范围查询物流订单详情 APIResponse
taobao.wlb.orderdetail.date.get

外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
*/
type TaobaoWlbOrderdetailDateGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbOrderdetailDateGetResponse `json:"taobao_wlb_orderdetail_date_get_response,omitempty"`
}

type TaobaoWlbOrderdetailDateGetResponse struct {

    // 物流宝订单，并且包含订单详情
    OrderDetailList   []WlbOrderDetail `json:"order_detail_list,omitempty"`

    // 总数
    TotalCount   int64 `json:"total_count,omitempty"`

}
