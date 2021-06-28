package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
大商家商品发布接口 APIResponse
tmall.item.vip.schema.add

大商家商品发布接口
*/
type TmallItemVipSchemaAddAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemVipSchemaAddResponse `json:"tmall_item_vip_schema_add_response,omitempty"` 
    TmallItemVipSchemaAddResponse
}

/* model for simplify = false
type TmallItemVipSchemaAddResponse struct {

    // 商品发布成功_商品id
    
    AddItemResult   string `json:"add_item_result,omitempty"`
    

    // sku与outerId映射信息
    
    SkuMapJson   string `json:"sku_map_json,omitempty"`
    

    // 发布商品操作成功时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

}
*/

type TmallItemVipSchemaAddResponse struct {

    // 商品发布成功_商品id
    AddItemResult   string `json:"add_item_result,omitempty"`

    // sku与outerId映射信息
    SkuMapJson   string `json:"sku_map_json,omitempty"`

    // 发布商品操作成功时间
    GmtCreate   string `json:"gmt_create,omitempty"`

}
