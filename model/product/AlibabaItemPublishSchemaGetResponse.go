package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品发布规则信息 API返回值 
alibaba.item.publish.schema.get

新商品发布，获取商品发布规则信息
*/
type AlibabaItemPublishSchemaGetAPIResponse struct {
    model.CommonResponse
    AlibabaItemPublishSchemaGetResponse
}

// 获取商品发布规则信息 成功返回结果
type AlibabaItemPublishSchemaGetResponse struct {
    XMLName xml.Name `xml:"alibaba_item_publish_schema_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商品发布规则信息，XML格式.
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
