package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查找IC商品或分销商品与后端商品的关联信息 APIResponse
taobao.scitem.map.query

查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
*/
type TaobaoScitemMapQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoScitemMapQueryResponse `json:"scitem_map_query_response,omitempty"` 
    TaobaoScitemMapQueryResponse
}

/* model for simplify = false
type TaobaoScitemMapQueryResponse struct {

    // 后端商品映射列表
    
    ScItemMaps  struct {
        ScItemMap  []ScItemMap `json:"sc_item_map,omitempty"`
    } `json:"sc_item_maps,omitempty"`
    

}
*/

type TaobaoScitemMapQueryResponse struct {

    // 后端商品映射列表
    ScItemMaps   []ScItemMap `json:"sc_item_maps,omitempty"`

}
