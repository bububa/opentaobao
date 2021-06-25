package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
vip商家发布商品的获取规则接口 APIResponse
tmall.item.vip.add.schema.get

获取vip商家发布商品的规则
*/
type TmallItemVipAddSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *TmallItemVipAddSchemaGetResponse `json:"tmall_item_vip_add_schema_get_response,omitempty"`
}

type TmallItemVipAddSchemaGetResponse struct {

    // 返回值是发布商品时需要的字段及基本类型
    Result   string `json:"result,omitempty"`

}
