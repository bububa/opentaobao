package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
编辑仓库覆盖范围 APIResponse
taobao.region.warehouse.manage

编辑仓库覆盖范围
*/
type TaobaoRegionWarehouseManageAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRegionWarehouseManageResponse `json:"region_warehouse_manage_response,omitempty"` 
    TaobaoRegionWarehouseManageResponse
}

/* model for simplify = false
type TaobaoRegionWarehouseManageResponse struct {

    // 返回结果
    
    Result  *struct {
        BaseResult  *BaseResult `json:"base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRegionWarehouseManageResponse struct {

    // 返回结果
    Result   *BaseResult `json:"result,omitempty"`

}
