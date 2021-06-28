package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ToB仓储发货 APIResponse
taobao.uop.tob.order.create

ToB仓储发货
*/
type TaobaoUopTobOrderCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"uop_tob_order_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // flag
    
    Flag   string `json:"flag,omitempty" xml:"