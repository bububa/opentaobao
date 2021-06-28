package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则编辑商品 APIResponse
tmall.item.schema.update

天猫根据规则编辑商品
*/
type TmallItemSchemaUpdateAPIResponse struct {
    model.CommonResponse
    TmallItemSchemaUpdateResponse
}

type TmallItemSchemaUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_item_schema_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回商品发布结果
    
    UpdateItemResult   string `json:"update_item_result,omitempty" xml:"update_item_result,omitempty"`

    
    // 商品更新操作成功时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`

    
}
