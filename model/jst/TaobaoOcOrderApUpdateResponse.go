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
    TaobaoOcOrderApUpdateResponse
}

type TaobaoOcOrderApUpdateResponse struct {
    XMLName xml.Name `xml:"oc_order_ap_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 描述操作执行是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // OC订单id
    
    OcOrderId   int64 `json:"oc_order_id,omitempty" xml:"oc_order_id,omitempty"`

    
    // 执行失败时候的错误描述信息
    
    ErrorDescription   string `json:"error_description,omitempty" xml:"error_description,omitempty"`

    
}
