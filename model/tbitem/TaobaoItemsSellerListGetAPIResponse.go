package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemssellerlistgetAPIResponse 批量获取商品详细信息 API返回值
// taobao.items.seller.list.get
//
// 批量获取商品详细信息
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoitemssellerlistgetAPIResponse struct {
	model.CommonResponse
	TaobaoitemssellerlistgetAPIResponseModel
}

// TaobaoitemssellerlistgetAPIResponseModel is 批量获取商品详细信息 成功返回结果
type TaobaoitemssellerlistgetAPIResponseModel struct {
	XMLName xml.Name `xml:"items_seller_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品详细信息列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
