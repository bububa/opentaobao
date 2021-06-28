package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建或更新仓库 APIResponse
taobao.inventory.store.manage

创建商家仓或者更新商家仓信息
*/
type TaobaoInventoryStoreManageAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoInventoryStoreManageResponse `json:"inventory_store_manage_response,omitempty"` 
    TaobaoInventoryStoreManageResponse
}

/* model for simplify = false
type TaobaoInventoryStoreManageResponse struct {

    // 返回结果
    
    StoreList  struct {
        Store  []Store `json:"store,omitempty"`
    } `json:"store_list,omitempty"`
    

}
*/

type TaobaoInventoryStoreManageResponse struct {

    // 返回结果
    StoreList   []Store `json:"store_list,omitempty"`

}
