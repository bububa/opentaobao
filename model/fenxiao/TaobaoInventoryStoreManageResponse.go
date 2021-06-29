package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建或更新仓库 API返回值 
taobao.inventory.store.manage

创建商家仓或者更新商家仓信息
*/
type TaobaoInventoryStoreManageAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryStoreManageResponse
}

// 创建或更新仓库 成功返回结果
type TaobaoInventoryStoreManageResponse struct {
    XMLName xml.Name `xml:"inventory_store_manage_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    StoreList   []Store `json:"store_list,omitempty" xml:"store_list>store,omitempty"`
}
