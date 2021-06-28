package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品仓库库存 APIResponse
tmall.inventory.query.forstore

商家查询后端商品仓库库存
*/
type TmallInventoryQueryForstoreAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_inventory_query_forstore_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果
    
    Result   *InventoryQueryResult `json:"result,omitempty" xml:"