package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商品发布规则信息 APIResponse
alibaba.item.publish.schema.get

新商品发布，获取商品发布规则信息
*/
type AlibabaItemPublishSchemaGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaItemPublishSchemaGetResponse `json:"alibaba_item_publish_schema_get_response,omitempty"`
}

type AlibabaItemPublishSchemaGetResponse struct {

    // 商品发布规则信息，XML格式.
    Result   string `json:"result,omitempty"`

}
