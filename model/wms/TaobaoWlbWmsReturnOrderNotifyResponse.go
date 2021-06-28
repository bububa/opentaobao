package wms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销售退货通知 APIResponse
taobao.wlb.wms.return.order.notify

销售退货通知
*/
type TaobaoWlbWmsReturnOrderNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWmsReturnOrderNotifyResponse
}

type TaobaoWlbWmsReturnOrderNotifyResponse struct {
    XMLName xml.Name `xml:"wlb_wms_return_order_notify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订单创建时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`

    
    // LBX929829111
    
    StoreOrderCode   string `json:"store_order_code,omitempty" xml:"store_order_code,omitempty"`

    
    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty" xml:"wl_success,omitempty"`

    
    // 错误编码
    
    WlErrorCode   string `json:"wl_error_code,omitempty" xml:"wl_error_code,omitempty"`

    
    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty" xml:"wl_error_msg,omitempty"`

    
}
