package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OC订单 APIResponse
taobao.oc.order.create

创建OC订单接口
*/
type TaobaoOcOrderCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"oc_order_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 表示是否执行成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"