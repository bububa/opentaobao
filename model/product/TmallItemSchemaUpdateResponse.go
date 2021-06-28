package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则编辑商品 APIResponse
tmall.item.schema.update

天猫根据规则编辑商品
*/
type TmallItemSchemaUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemSchemaUpdateResponse `json:"tmall_item_schema_update_response,omitempty"` 
    TmallItemSchemaUpdateResponse
}

/* model for simplify = false
type TmallItemSchemaUpdateResponse struct {

    // 返回商品发布结果
    
    UpdateItemResult   string `json:"update_item_result,omitempty"`
    

    // 商品更新操作成功时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

}
*/

type TmallItemSchemaUpdateResponse struct {

    // 返回商品发布结果
    UpdateItemResult   string `json:"update_item_result,omitempty"`

    // 商品更新操作成功时间
    GmtModified   string `json:"gmt_modified,omitempty"`

}
