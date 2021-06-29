package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家可发布商品的市场信息 API返回值 
alibaba.item.publish.market.get

获取商家可发布商品的市场信息
*/
type AlibabaItemPublishMarketGetAPIResponse struct {
    model.CommonResponse
    AlibabaItemPublishMarketGetResponse
}

// 获取商家可发布商品的市场信息 成功返回结果
type AlibabaItemPublishMarketGetResponse struct {
    XMLName xml.Name `xml:"alibaba_item_publish_market_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商家可发布的市场列表,多个以逗号(,)分隔
    Markets   string `json:"markets,omitempty" xml:"markets,omitempty"`
}
