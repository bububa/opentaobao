package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库信息 APIResponse
taobao.inventory.store.query

查询商家仓信息
*/
type TaobaoInventoryStoreQueryAPIResponse struct {
    model.CommonResponse
    TaobaoInventoryStoreQueryResponse
}

type TaobaoInventoryStoreQueryResponse struct {
    XMLName xml.Name `xml:"inventory_store_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 仓库列表
    
    StoreList   []Store `json:"store_list,omitempty" xml:"store_list>store,omitempty"`
    
    
}
