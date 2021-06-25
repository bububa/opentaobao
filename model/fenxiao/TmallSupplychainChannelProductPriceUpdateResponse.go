package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
渠道价格更新接口 APIResponse
tmall.supplychain.channel.product.price.update

更新渠道产品价格
*/
type TmallSupplychainChannelProductPriceUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallSupplychainChannelProductPriceUpdateResponse `json:"tmall_supplychain_channel_product_price_update_response,omitempty"`
}

type TmallSupplychainChannelProductPriceUpdateResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *ResultDTO `json:"result,omitempty"`

}
