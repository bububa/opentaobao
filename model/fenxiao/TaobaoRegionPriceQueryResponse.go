package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
区域价格查询 APIResponse
taobao.region.price.query

区域价格查询
*/
type TaobaoRegionPriceQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRegionPriceQueryResponse `json:"taobao_region_price_query_response,omitempty"`
}

type TaobaoRegionPriceQueryResponse struct {

    // result
    Result   *BaseResult `json:"result,omitempty"`

}
