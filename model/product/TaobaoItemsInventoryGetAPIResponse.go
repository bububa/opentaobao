package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
得到当前会话用户库存中的商品列表 API返回值 
taobao.items.inventory.get

获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤 
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取<br/>
<strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemsInventoryGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemsInventoryGetAPIResponseModel
}

// 得到当前会话用户库存中的商品列表 成功返回结果
type TaobaoItemsInventoryGetAPIResponseModel struct {
    XMLName xml.Name `xml:"items_inventory_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 搜索到底商品列表，具体字段根据设定的fields决定，不包括desc,stuff_status字段
    Items   []Item `json:"items,omitempty" xml:"items>item,omitempty"`
    // 搜索到符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
