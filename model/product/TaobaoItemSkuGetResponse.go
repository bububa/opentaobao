package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取SKU APIResponse
taobao.item.sku.get

获取sku_id所对应的sku数据 
sku_id对应的sku要属于传入的nick对应的卖家
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
*/
type TaobaoItemSkuGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemSkuGetResponse
}

type TaobaoItemSkuGetResponse struct {
    XMLName xml.Name `xml:"item_sku_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Sku
    
    Sku   *Sku `json:"sku,omitempty" xml:"sku,omitempty"`

    
}
