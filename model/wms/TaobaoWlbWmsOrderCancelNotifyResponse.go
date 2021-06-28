package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单据取消接口 APIResponse
taobao.wlb.wms.order.cancel.notify

单据取消接口
*/
type TaobaoWlbWmsOrderCancelNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_wms_order_cancel_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"