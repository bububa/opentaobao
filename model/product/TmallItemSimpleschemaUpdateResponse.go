package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫简化编辑商品 APIResponse
tmall.item.simpleschema.update

国外大商家天猫简化编辑商品
*/
type TmallItemSimpleschemaUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallItemSimpleschemaUpdateResponse `json:"tmall_item_simpleschema_update_response,omitempty"`
}

type TmallItemSimpleschemaUpdateResponse struct {

    // sku与outerId映射信息
    SkuMapJson   string `json:"sku_map_json,omitempty"`

    // 编辑商品的itemid
    UpdateItemResult   string `json:"update_item_result,omitempty"`

    // 编辑商品操作成功时间
    GmtModified   string `json:"gmt_modified,omitempty"`

}
