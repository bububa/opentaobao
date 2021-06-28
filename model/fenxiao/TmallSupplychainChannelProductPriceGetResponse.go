package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道价格查询接口 APIResponse
tmall.supplychain.channel.product.price.get

渠道价格查询接口
*/
type TmallSupplychainChannelProductPriceGetAPIResponse struct {
    model.CommonResponse
    TmallSupplychainChannelProductPriceGetResponse
}

type TmallSupplychainChannelProductPriceGetResponse struct {
    XMLName xml.Name `xml:"tmall_supplychain_channel_product_price_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *TmallSupplychainChannelProductPriceGetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
