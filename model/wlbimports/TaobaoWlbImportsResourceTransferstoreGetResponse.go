package wlbimports

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据指定的资源获取所有中转仓列表 APIResponse
taobao.wlb.imports.resource.transferstore.get

根据指定的资源获取所有中转仓列表
*/
type TaobaoWlbImportsResourceTransferstoreGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbImportsResourceTransferstoreGetResponse `json:"wlb_imports_resource_transferstore_get_response,omitempty"` 
    TaobaoWlbImportsResourceTransferstoreGetResponse
}

/* model for simplify = false
type TaobaoWlbImportsResourceTransferstoreGetResponse struct {

    // 符合条件的中转仓列表
    
    Stores  struct {
        TranStoreResult  []TranStoreResult `json:"tran_store_result,omitempty"`
    } `json:"stores,omitempty"`
    

}
*/

type TaobaoWlbImportsResourceTransferstoreGetResponse struct {

    // 符合条件的中转仓列表
    Stores   []TranStoreResult `json:"stores,omitempty"`

}
