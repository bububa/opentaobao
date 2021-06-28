package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
失效指定用户的商品与后端商品的映射关系 APIResponse
taobao.scitem.map.delete

根据后端商品Id，失效指定用户的商品与后端商品的映射关系
*/
type TaobaoScitemMapDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoScitemMapDeleteResponse `json:"scitem_map_delete_response,omitempty"` 
    TaobaoScitemMapDeleteResponse
}

/* model for simplify = false
type TaobaoScitemMapDeleteResponse struct {

    // 失效条数
    
    Module   int64 `json:"module,omitempty"`
    

}
*/

type TaobaoScitemMapDeleteResponse struct {

    // 失效条数
    Module   int64 `json:"module,omitempty"`

}
