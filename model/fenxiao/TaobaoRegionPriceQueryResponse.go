package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
区域价格查询 APIResponse
taobao.region.price.query

区域价格查询
*/
type TaobaoRegionPriceQueryAPIResponse struct {
    model.CommonResponse
    TaobaoRegionPriceQueryResponse
}

type TaobaoRegionPriceQueryResponse struct {
    XMLName xml.Name `xml:"region_price_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
