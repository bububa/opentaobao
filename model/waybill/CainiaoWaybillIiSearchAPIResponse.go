package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询面单服务订购及面单使用情况 API返回值 
cainiao.waybill.ii.search

获取发货地&CP开通状态&账户的使用情况
*/
type CainiaoWaybillIiSearchAPIResponse struct {
    model.CommonResponse
    CainiaoWaybillIiSearchAPIResponseModel
}

// 查询面单服务订购及面单使用情况 成功返回结果
type CainiaoWaybillIiSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_waybill_ii_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // CP网点信息及对应的商家的发货信息
    WaybillApplySubscriptionCols   []WaybillApplySubscriptionInfo `json:"waybill_apply_subscription_cols,omitempty" xml:"waybill_apply_subscription_cols>waybill_apply_subscription_info,omitempty"`
}
