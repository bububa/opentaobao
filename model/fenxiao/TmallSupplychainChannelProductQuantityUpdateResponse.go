package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
渠道无仓库存更新接口 APIResponse
tmall.supplychain.channel.product.quantity.update

渠道无仓库存更新接口
*/
type TmallSupplychainChannelProductQuantityUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallSupplychainChannelProductQuantityUpdateResponse `json:"tmall_supplychain_channel_product_quantity_update_response,omitempty"` 
    TmallSupplychainChannelProductQuantityUpdateResponse
}

/* model for simplify = false
type TmallSupplychainChannelProductQuantityUpdateResponse struct {

    // 异步获取历史数据接口返回结果
    
    Result  *struct {
        TmallSupplychainChannelProductQuantityUpdateResultDto  *TmallSupplychainChannelProductQuantityUpdateResultDto `json:"tmall_supplychain_channel_product_quantity_update_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallSupplychainChannelProductQuantityUpdateResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *TmallSupplychainChannelProductQuantityUpdateResultDto `json:"result,omitempty"`

}
