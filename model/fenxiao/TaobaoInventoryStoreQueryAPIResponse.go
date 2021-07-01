package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库信息 API返回值 
taobao.inventory.store.query

查询商家仓信息
*/
type TaobaoInventoryStoreQueryAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryStoreQueryAPIResponseModel
}

// 查询仓库信息 成功返回结果
type TaobaoInventoryStoreQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"inventory_store_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 仓库列表
    StoreList   []Store `json:"store_list,omitempty" xml:"store_list>store,omitempty"`
}
