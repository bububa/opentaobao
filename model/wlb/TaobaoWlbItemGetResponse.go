package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID获取商品信息 APIResponse
taobao.wlb.item.get

根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
*/
type TaobaoWlbItemGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbItemGetResponse
}

type TaobaoWlbItemGetResponse struct {
    XMLName xml.Name `xml:"wlb_item_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品信息
    
    Item   *WlbItem `json:"item,omitempty" xml:"item,omitempty"`

    
}
