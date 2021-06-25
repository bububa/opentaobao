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
    Response *TaobaoWlbImportsResourceTransferstoreGetResponse `json:"taobao_wlb_imports_resource_transferstore_get_response,omitempty"`
}

type TaobaoWlbImportsResourceTransferstoreGetResponse struct {

    // 符合条件的中转仓列表
    Stores   []TranStoreResult `json:"stores,omitempty"`

}
