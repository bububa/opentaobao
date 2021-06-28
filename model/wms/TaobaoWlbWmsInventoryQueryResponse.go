package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟商品库存查询 APIResponse
taobao.wlb.wms.inventory.query

支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
*/
type TaobaoWlbWmsInventoryQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsInventoryQueryResponse `json:"wlb_wms_inventory_query_response,omitempty"` 
    TaobaoWlbWmsInventoryQueryResponse
}

/* model for simplify = false
type TaobaoWlbWmsInventoryQueryResponse struct {

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

    // 错误代码
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

    // 总数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 商品详情列表
    
    ItemList  struct {
        WmsInventoryQueryItemlist  []WmsInventoryQueryItemlist `json:"wms_inventory_query_itemlist,omitempty"`
    } `json:"item_list,omitempty"`
    

}
*/

type TaobaoWlbWmsInventoryQueryResponse struct {

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误代码
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 错误信息
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 总数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 商品详情列表
    ItemList   []WmsInventoryQueryItemlist `json:"item_list,omitempty"`

}
