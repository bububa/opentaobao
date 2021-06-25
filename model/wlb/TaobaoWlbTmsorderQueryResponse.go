package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过物流订单编号查询物流信息 APIResponse
taobao.wlb.tmsorder.query

通过物流订单编号分页查询物流信息
*/
type TaobaoWlbTmsorderQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbTmsorderQueryResponse `json:"taobao_wlb_tmsorder_query_response,omitempty"`
}

type TaobaoWlbTmsorderQueryResponse struct {

    // 物流订单运单信息列表
    TmsOrderList   []WlbTmsOrder `json:"tms_order_list,omitempty"`

    // 结果总数
    TotalCount   int64 `json:"total_count,omitempty"`

}