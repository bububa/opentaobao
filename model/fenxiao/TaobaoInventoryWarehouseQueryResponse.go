package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询商家仓信息 APIResponse
taobao.inventory.warehouse.query

分页查询商家仓信息
*/
type TaobaoInventoryWarehouseQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"inventory_warehouse_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *PaginationResult `json:"result,omitempty" xml:"