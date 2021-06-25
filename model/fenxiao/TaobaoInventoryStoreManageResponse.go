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
    Response *TaobaoInventoryStoreManageResponse `json:"taobao_inventory_store_manage_response,omitempty"`
}

type TaobaoInventoryStoreManageResponse struct {

    // 返回结果
    StoreList   []Store `json:"store_list,omitempty"`

}
