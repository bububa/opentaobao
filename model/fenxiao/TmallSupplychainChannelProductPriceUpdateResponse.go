package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道价格更新接口 APIResponse
tmall.supplychain.channel.product.price.update

更新渠道产品价格
*/
type TmallSupplychainChannelProductPriceUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_supplychain_channel_product_price_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *ResultDTO `json:"result,omitempty" xml:"