package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据物流宝订单编号查询物流宝订单概要信息 APIResponse
taobao.wlb.wlborder.get

根据物流宝订单编号查询物流宝订单概要信息
*/
type TaobaoWlbWlborderGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbWlborderGetResponse
}

type TaobaoWlbWlborderGetResponse struct {
    XMLName xml.Name `xml:"wlb_wlborder_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 物流宝订单对象
    
    WlbOrder   *WlbOrder `json:"wlb_order,omitempty" xml:"wlb_order,omitempty"`

    
}
