package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID查询所有库存变更记录 APIResponse
taobao.wlb.inventorylog.query

通过商品ID等几个条件来分页查询库存变更记录
*/
type TaobaoWlbInventorylogQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbInventorylogQueryResponse `json:"taobao_wlb_inventorylog_query_response,omitempty"`
}

type TaobaoWlbInventorylogQueryResponse struct {

    // 记录数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 库存变更记录
    InventoryLogList   []WlbItemInventoryLog `json:"inventory_log_list,omitempty"`

}
