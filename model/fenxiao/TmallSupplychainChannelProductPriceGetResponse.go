package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道价格查询接口 API返回值 
tmall.supplychain.channel.product.price.get

渠道价格查询接口
*/
type TmallSupplychainChannelProductPriceGetAPIResponse struct {
    model.CommonResponse
    TmallSupplychainChannelProductPriceGetResponse
}

// 渠道价格查询接口 成功返回结果
type TmallSupplychainChannelProductPriceGetResponse struct {
    XMLName xml.Name `xml:"tmall_supplychain_channel_product_price_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步获取历史数据接口返回结果
    Result   *TmallSupplychainChannelProductPriceGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
