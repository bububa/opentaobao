package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取商品详细信息 APIResponse
taobao.items.seller.list.get

批量获取商品详细信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemsSellerListGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemsSellerListGetResponse
}

type TaobaoItemsSellerListGetResponse struct {
    XMLName xml.Name `xml:"items_seller_list_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品详细信息列表
    
    Items   []Item `json:"items,omitempty" xml:"items>item,omitempty"`
    
    
}
