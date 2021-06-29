package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道中心-查询产品列表 APIResponse
tmall.channel.products.query

渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
*/
type TmallChannelProductsQueryAPIResponse struct {
    model.CommonResponse
    TmallChannelProductsQueryResponse
}

type TmallChannelProductsQueryResponse struct {
    XMLName xml.Name `xml:"tmall_channel_products_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PageResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
