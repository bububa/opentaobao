package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
供应商铺货 APIResponse
tmall.supplychain.channel.product.release

供应商渠道铺货接口
*/
type TmallSupplychainChannelProductReleaseAPIResponse struct {
    model.CommonResponse
    // Response *TmallSupplychainChannelProductReleaseResponse `json:"tmall_supplychain_channel_product_release_response,omitempty"` 
    TmallSupplychainChannelProductReleaseResponse
}

/* model for simplify = false
type TmallSupplychainChannelProductReleaseResponse struct {

    // 返回结果
    
    Result  *struct {
        TmallSupplychainChannelProductReleaseResultDto  *TmallSupplychainChannelProductReleaseResultDto `json:"tmall_supplychain_channel_product_release_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallSupplychainChannelProductReleaseResponse struct {

    // 返回结果
    Result   *TmallSupplychainChannelProductReleaseResultDto `json:"result,omitempty"`

}
