package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品仓库库存 APIResponse
tmall.inventory.query.forstore

商家查询后端商品仓库库存
*/
type TmallInventoryQueryForstoreAPIResponse struct {
    model.CommonResponse
    TmallInventoryQueryForstoreResponse
}

type TmallInventoryQueryForstoreResponse struct {
    XMLName xml.Name `xml:"tmall_inventory_query_forstore_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果
    
    Result   *InventoryQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
    // 错误code
    
    Errorcode   string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`

    
    // 错误信息
    
    Errormessage   string `json:"errormessage,omitempty" xml:"errormessage,omitempty"`

    
    // 整体成功或失败
    
    Issuccess   bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`

    
}
