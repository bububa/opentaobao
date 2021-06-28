package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询库存详情 APIResponse
taobao.wlb.inventory.detail.get

查询库存详情，通过商品ID获取发送请求的卖家的库存详情
*/
type TaobaoWlbInventoryDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbInventoryDetailGetResponse `json:"wlb_inventory_detail_get_response,omitempty"` 
    TaobaoWlbInventoryDetailGetResponse
}

/* model for simplify = false
type TaobaoWlbInventoryDetailGetResponse struct {

    // 库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性
    
    InventoryList  struct {
        WlbInventory  []WlbInventory `json:"wlb_inventory,omitempty"`
    } `json:"inventory_list,omitempty"`
    

    // 入参的item_id
    
    ItemId   int64 `json:"item_id,omitempty"`
    

}
*/

type TaobaoWlbInventoryDetailGetResponse struct {

    // 库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性
    InventoryList   []WlbInventory `json:"inventory_list,omitempty"`

    // 入参的item_id
    ItemId   int64 `json:"item_id,omitempty"`

}
