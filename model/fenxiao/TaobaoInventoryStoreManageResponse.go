package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建或更新仓库 APIResponse
taobao.inventory.store.manage

创建商家仓或者更新商家仓信息
*/
type TaobaoInventoryStoreManageAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"inventory_store_manage_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    StoreList   []Store `json:"store_list,omitempty" xml:"