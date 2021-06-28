package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
大商家商品编辑接口 APIResponse
tmall.item.vip.schema.update

大商家编辑商品
*/
type TmallItemVipSchemaUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemVipSchemaUpdateResponse `json:"tmall_item_vip_schema_update_response,omitempty"` 
    TmallItemVipSchemaUpdateResponse
}

/* model for simplify = false
type TmallItemVipSchemaUpdateResponse struct {

    // 编辑商品的id
    
    UpdateItemResult   string `json:"update_item_result,omitempty"`
    

    // sku与outerId映射信息
    
    SkuMapJson   string `json:"sku_map_json,omitempty"`
    

    // 编辑商品操作成功时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

}
*/

type TmallItemVipSchemaUpdateResponse struct {

    // 编辑商品的id
    UpdateItemResult   string `json:"update_item_result,omitempty"`

    // sku与outerId映射信息
    SkuMapJson   string `json:"sku_map_json,omitempty"`

    // 编辑商品操作成功时间
    GmtCreate   string `json:"gmt_create,omitempty"`

}
