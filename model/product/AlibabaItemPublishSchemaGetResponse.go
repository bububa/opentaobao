package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品发布规则信息 APIResponse
alibaba.item.publish.schema.get

新商品发布，获取商品发布规则信息
*/
type AlibabaItemPublishSchemaGetAPIResponse struct {
    model.CommonResponse
    AlibabaItemPublishSchemaGetResponse
}

type AlibabaItemPublishSchemaGetResponse struct {
    XMLName xml.Name `xml:"alibaba_item_publish_schema_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品发布规则信息，XML格式.
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
