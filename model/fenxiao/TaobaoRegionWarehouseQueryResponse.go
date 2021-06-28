package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库覆盖范围 APIResponse
taobao.region.warehouse.query

查询仓库覆盖范围
*/
type TaobaoRegionWarehouseQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRegionWarehouseQueryResponse `json:"region_warehouse_query_response,omitempty"` 
    TaobaoRegionWarehouseQueryResponse
}

/* model for simplify = false
type TaobaoRegionWarehouseQueryResponse struct {

    // result
    
    Result  *struct {
        BaseResult  *BaseResult `json:"base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRegionWarehouseQueryResponse struct {

    // result
    Result   *BaseResult `json:"result,omitempty"`

}
