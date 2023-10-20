package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSkuGetAPIResponse 获取SKU API返回值
// taobao.item.sku.get
//
// 获取sku_id所对应的sku数据
// sku_id对应的sku要属于传入的nick对应的卖家
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoItemSkuGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemSkuGetAPIResponseModel
}

// TaobaoItemSkuGetAPIResponseModel is 获取SKU 成功返回结果
type TaobaoItemSkuGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_sku_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Sku
	Sku *Sku `json:"sku,omitempty" xml:"sku,omitempty"`
}
