package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询面单服务订购及面单使用情况 APIResponse
cainiao.waybill.ii.search

获取发货地&CP开通状态&账户的使用情况
*/
type CainiaoWaybillIiSearchAPIResponse struct {
    model.CommonResponse
    Response *CainiaoWaybillIiSearchResponse `json:"cainiao_waybill_ii_search_response,omitempty"`
}

type CainiaoWaybillIiSearchResponse struct {

    // CP网点信息及对应的商家的发货信息
    WaybillApplySubscriptionCols   []WaybillApplySubscriptionInfo `json:"waybill_apply_subscription_cols,omitempty"`

}