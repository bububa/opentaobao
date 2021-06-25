package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
渠道库存查询接口 APIResponse
tmall.supplychain.channel.product.quantity.get

渠道库存查询接口
*/
type TmallSupplychainChannelProductQuantityGetAPIResponse struct {
    model.CommonResponse
    Response *TmallSupplychainChannelProductQuantityGetResponse `json:"tmall_supplychain_channel_product_quantity_get_response,omitempty"`
}

type TmallSupplychainChannelProductQuantityGetResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *TmallSupplychainChannelProductQuantityGetResultDto `json:"result,omitempty"`

}
