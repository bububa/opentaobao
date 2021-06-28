package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品下架 APIResponse
tmall.supplychain.channel.product.downshelf

产品下架
*/
type TmallSupplychainChannelProductDownshelfAPIResponse struct {
    model.CommonResponse
    TmallSupplychainChannelProductDownshelfResponse
}

type TmallSupplychainChannelProductDownshelfResponse struct {
    XMLName xml.Name `xml:"tmall_supplychain_channel_product_downshelf_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 异步获取历史数据接口返回结果
    
    Result   *TmallSupplychainChannelProductDownshelfResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
