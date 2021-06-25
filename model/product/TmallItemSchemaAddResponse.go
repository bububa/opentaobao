package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则发布商品 APIResponse
tmall.item.schema.add

天猫TopSchema发布商品。
*/
type TmallItemSchemaAddAPIResponse struct {
    model.CommonResponse
    Response *TmallItemSchemaAddResponse `json:"tmall_item_schema_add_response,omitempty"`
}

type TmallItemSchemaAddResponse struct {

    // 返回商品发布结果
    AddItemResult   string `json:"add_item_result,omitempty"`

    // 发布商品操作成功时间
    GmtCreate   string `json:"gmt_create,omitempty"`

}
