package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品级联属性信息获取 APIResponse
alibaba.item.publish.props.get

新商品发布，商品级联属性信息获取
*/
type AlibabaItemPublishPropsGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaItemPublishPropsGetResponse `json:"alibaba_item_publish_props_get_response,omitempty"`
}

type AlibabaItemPublishPropsGetResponse struct {

    // 商品发布规则信息，XML格式.
    Result   string `json:"result,omitempty"`

}
