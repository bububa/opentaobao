package wlbimports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据指定的资源获取所有中转仓列表 APIResponse
taobao.wlb.imports.resource.transferstore.get

根据指定的资源获取所有中转仓列表
*/
type TaobaoWlbImportsResourceTransferstoreGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_imports_resource_transferstore_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 符合条件的中转仓列表
    
    Stores   []TranStoreResult `json:"stores,omitempty" xml:"