package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道库存查询接口 APIResponse
tmall.supplychain.channel.product.quantity.get

渠道库存查询接口
*/
type TmallSupplychainChannelProductQuantityGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_supplychain_channel_product_quantity_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *TmallSupplychainChannelProductQuantityGetResultDto `json:"result,omitempty" xml:"