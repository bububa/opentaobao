package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品上架 APIResponse
tmall.supplychain.channel.product.upshelf

上架渠道产品
*/
type TmallSupplychainChannelProductUpshelfAPIResponse struct {
    model.CommonResponse
    // Response *TmallSupplychainChannelProductUpshelfResponse `json:"tmall_supplychain_channel_product_upshelf_response,omitempty"` 
    TmallSupplychainChannelProductUpshelfResponse
}

/* model for simplify = false
type TmallSupplychainChannelProductUpshelfResponse struct {

    // 异步获取历史数据接口返回结果
    
    Result  *struct {
        TmallSupplychainChannelProductUpshelfResultDto  *TmallSupplychainChannelProductUpshelfResultDto `json:"tmall_supplychain_channel_product_upshelf_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallSupplychainChannelProductUpshelfResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *TmallSupplychainChannelProductUpshelfResultDto `json:"result,omitempty"`

}
