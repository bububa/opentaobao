package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫简化编辑商品 APIResponse
tmall.item.simpleschema.update

国外大商家天猫简化编辑商品
*/
type TmallItemSimpleschemaUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_simpleschema_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // sku与outerId映射信息
    
    SkuMapJson   string `json:"sku_map_json,omitempty" xml:"