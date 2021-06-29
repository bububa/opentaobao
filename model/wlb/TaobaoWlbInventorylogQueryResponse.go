package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据商品ID查询所有库存变更记录 APIResponse
taobao.wlb.inventorylog.query

通过商品ID等几个条件来分页查询库存变更记录
*/
type TaobaoWlbInventorylogQueryAPIResponse struct {
    model.CommonResponse
    TaobaoWlbInventorylogQueryResponse
}

type TaobaoWlbInventorylogQueryResponse struct {
    XMLName xml.Name `xml:"wlb_inventorylog_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 记录数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
    // 库存变更记录
    
    InventoryLogList   []WlbItemInventoryLog `json:"inventory_log_list,omitempty" xml:"inventory_log_list>wlb_item_inventory_log,omitempty"`
    
    
}
