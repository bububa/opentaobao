package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品上架 APIResponse
tmall.supplychain.channel.product.upshelf

上架渠道产品
*/
type TmallSupplychainChannelProductUpshelfAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_supplychain_channel_product_upshelf_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *TmallSupplychainChannelProductUpshelfResultDto `json:"result,omitempty" xml:"