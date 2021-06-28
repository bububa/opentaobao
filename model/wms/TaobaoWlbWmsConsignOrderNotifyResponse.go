package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货订单通知 APIResponse
taobao.wlb.wms.consign.order.notify

发货订单通知
*/
type TaobaoWlbWmsConsignOrderNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_consign_order_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"