package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消物流宝订单 APIResponse
taobao.wlb.order.cancel

取消物流宝订单
*/
type TaobaoWlbOrderCancelAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderCancelResponse
}

type TaobaoWlbOrderCancelResponse struct {
    XMLName xml.Name `xml:"wlb_order_cancel_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改时间，只有在取消成功的情况下，才可以做
    
    ModifyTime   string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`

    
    // 错误编码列表
    
    ErrorCodeList   string `json:"error_code_list,omitempty" xml:"error_code_list,omitempty"`

    
}
