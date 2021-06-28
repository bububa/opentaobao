package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑获取schema信息 APIResponse
alibaba.item.edit.schema.get

商品编辑时，获取商品规则信息
*/
type AlibabaItemEditSchemaGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaItemEditSchemaGetResponse `json:"alibaba_item_edit_schema_get_response,omitempty"` 
    AlibabaItemEditSchemaGetResponse
}

/* model for simplify = false
type AlibabaItemEditSchemaGetResponse struct {

    // 商品已有规则信息，XML格式.
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaItemEditSchemaGetResponse struct {

    // 商品已有规则信息，XML格式.
    Result   string `json:"result,omitempty"`

}
