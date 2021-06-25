package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询面单服务订购及面单使用情况v1.0 APIResponse
taobao.wlb.waybill.i.search

获取发货地&CP开通状态&账户的使用情况
*/
type TaobaoWlbWaybillISearchAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbWaybillISearchResponse `json:"taobao_wlb_waybill_i_search_response,omitempty"`
}

type TaobaoWlbWaybillISearchResponse struct {

    // 订购关系
    Subscribtions   []WaybillApplySubscriptionInfo `json:"subscribtions,omitempty"`

}
