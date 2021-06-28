package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建商家仓或者更新商家仓信息 APIResponse
taobao.inventory.warehouse.manage

创建商家仓或者更新商家仓信息
*/
type TaobaoInventoryWarehouseManageAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"inventory_warehouse_manage_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoInventoryWarehouseManageResult `json:"result,omitempty" xml:"