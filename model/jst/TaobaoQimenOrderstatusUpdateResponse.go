package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态更新接口 APIResponse
taobao.qimen.orderstatus.update

星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
*/
type TaobaoQimenOrderstatusUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_orderstatus_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // message
    
    Message   string `json:"message,omitempty" xml:"