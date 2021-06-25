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
    Response *TaobaoRegionWarehouseManageResponse `json:"taobao_region_warehouse_manage_response,omitempty"`
}

type TaobaoRegionWarehouseManageResponse struct {

    // 返回结果
    Result   *BaseResult `json:"result,omitempty"`

}
