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
    TaobaoWlbWmsOrderCancelNotifyResponse
}

type TaobaoWlbWmsOrderCancelNotifyResponse struct {
    XMLName xml.Name `xml:"wlb_wms_order_cancel_notify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`

    
    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`

    
    // 错误详细
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`

    
}
