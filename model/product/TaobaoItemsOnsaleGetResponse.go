package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取当前会话用户出售中的商品列表 APIResponse
taobao.items.onsale.get

获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get 获取
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemsOnsaleGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"items_onsale_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 搜索到的商品列表，具体字段根据设定的fields决定，不包括desc字段
    
    Items   []Item `json:"items,omitempty" xml:"