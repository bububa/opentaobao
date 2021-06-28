package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询面单服务订购及面单使用情况v1.0 APIResponse
taobao.wlb.waybill.i.search

获取发货地&CP开通状态&账户的使用情况
*/
type TaobaoWlbWaybillISearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_waybill_i_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订购关系
    
    Subscribtions   []WaybillApplySubscriptionInfo `json:"subscribtions,omitempty" xml:"