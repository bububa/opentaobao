package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库信息 APIResponse
taobao.inventory.store.query

查询商家仓信息
*/
type TaobaoInventoryStoreQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoInventoryStoreQueryResponse `json:"taobao_inventory_store_query_response,omitempty"`
}

type TaobaoInventoryStoreQueryResponse struct {

    // 仓库列表
    StoreList   []Store `json:"store_list,omitempty"`

}
