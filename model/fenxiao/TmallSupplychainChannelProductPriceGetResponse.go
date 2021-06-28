package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
渠道价格查询接口 APIResponse
tmall.supplychain.channel.product.price.get

渠道价格查询接口
*/
type TmallSupplychainChannelProductPriceGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallSupplychainChannelProductPriceGetResponse `json:"tmall_supplychain_channel_product_price_get_response,omitempty"` 
    TmallSupplychainChannelProductPriceGetResponse
}

/* model for simplify = false
type TmallSupplychainChannelProductPriceGetResponse struct {

    // 异步获取历史数据接口返回结果
    
    Result  *struct {
        TmallSupplychainChannelProductPriceGetResultDto  *TmallSupplychainChannelProductPriceGetResultDto `json:"tmall_supplychain_channel_product_price_get_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallSupplychainChannelProductPriceGetResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *TmallSupplychainChannelProductPriceGetResultDto `json:"result,omitempty"`

}
