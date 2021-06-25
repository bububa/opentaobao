package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品下架 APIResponse
tmall.supplychain.channel.product.downshelf

产品下架
*/
type TmallSupplychainChannelProductDownshelfAPIResponse struct {
    model.CommonResponse
    Response *TmallSupplychainChannelProductDownshelfResponse `json:"tmall_supplychain_channel_product_downshelf_response,omitempty"`
}

type TmallSupplychainChannelProductDownshelfResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *TmallSupplychainChannelProductDownshelfResultDto `json:"result,omitempty"`

}
