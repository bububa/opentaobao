package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑获取schema信息 APIResponse
alibaba.item.edit.schema.get

商品编辑时，获取商品规则信息
*/
type AlibabaItemEditSchemaGetAPIResponse struct {
    model.CommonResponse
    AlibabaItemEditSchemaGetResponse
}

type AlibabaItemEditSchemaGetResponse struct {
    XMLName xml.Name `xml:"alibaba_item_edit_schema_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品已有规则信息，XML格式.
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
