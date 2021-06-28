package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建物流宝订单 APIResponse
taobao.wlb.order.create

创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
*/
type TaobaoWlbOrderCreateAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderCreateResponse
}

type TaobaoWlbOrderCreateResponse struct {
    XMLName xml.Name `xml:"wlb_order_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 物流宝订单创建成功后，返回物流宝的订单编号；如果订单创建失败，该字段为空。
    
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`

    
    // 订单创建时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`

    
}
