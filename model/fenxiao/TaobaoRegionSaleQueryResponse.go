package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品销售区域 APIResponse
taobao.region.sale.query

查询商品销售区域
*/
type TaobaoRegionSaleQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRegionSaleQueryResponse `json:"region_sale_query_response,omitempty"` 
    TaobaoRegionSaleQueryResponse
}

/* model for simplify = false
type TaobaoRegionSaleQueryResponse struct {

    // result
    
    Result  *struct {
        BaseResult  *BaseResult `json:"base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRegionSaleQueryResponse struct {

    // result
    Result   *BaseResult `json:"result,omitempty"`

}
