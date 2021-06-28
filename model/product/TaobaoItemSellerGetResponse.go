package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个商品详细信息 APIResponse
taobao.item.seller.get

获取单个商品的全部信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemSellerGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemSellerGetResponse
}

type TaobaoItemSellerGetResponse struct {
    XMLName xml.Name `xml:"item_seller_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品详细信息
    
    Item   *Item `json:"item,omitempty" xml:"item,omitempty"`

    
}
