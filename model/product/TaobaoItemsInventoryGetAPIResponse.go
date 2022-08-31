package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemsInventoryGetAPIResponse 得到当前会话用户库存中的商品列表 API返回值
// taobao.items.inventory.get
//
// 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤
// 只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取&lt;br/&gt;
// &lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoItemsInventoryGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemsInventoryGetAPIResponseModel
}

// TaobaoItemsInventoryGetAPIResponseModel is 得到当前会话用户库存中的商品列表 成功返回结果
type TaobaoItemsInventoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"items_inventory_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到底商品列表，具体字段根据设定的fields决定，不包括desc,stuff_status字段
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 搜索到符合条件的结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
