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
    Response *TaobaoRegionSaleQueryResponse `json:"taobao_region_sale_query_response,omitempty"`
}

type TaobaoRegionSaleQueryResponse struct {

    // result
    Result   *BaseResult `json:"result,omitempty"`

}
