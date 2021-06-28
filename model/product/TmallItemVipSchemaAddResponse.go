package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
大商家商品发布接口 APIResponse
tmall.item.vip.schema.add

大商家商品发布接口
*/
type TmallItemVipSchemaAddAPIResponse struct {
    model.CommonResponse
    TmallItemVipSchemaAddResponse
}

type TmallItemVipSchemaAddResponse struct {
    XMLName xml.Name `xml:"tmall_item_vip_schema_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品发布成功_商品id
    
    AddItemResult   string `json:"add_item_result,omitempty" xml:"add_item_result,omitempty"`

    
    // sku与outerId映射信息
    
    SkuMapJson   string `json:"sku_map_json,omitempty" xml:"sku_map_json,omitempty"`

    
    // 发布商品操作成功时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`

    
}
