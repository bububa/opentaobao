package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按OC订单分账 APIResponse
taobao.oc.order.ap.update

对OC订单执行分账操作
*/
type TaobaoOcOrderApUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"oc_order_ap_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 描述操作执行是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"