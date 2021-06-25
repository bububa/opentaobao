package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品发布 APIResponse
alibaba.item.publish.submit

新商品发布，提交商品发布信息
*/
type AlibabaItemPublishSubmitAPIResponse struct {
    model.CommonResponse
    Response *AlibabaItemPublishSubmitResponse `json:"alibaba_item_publish_submit_response,omitempty"`
}

type AlibabaItemPublishSubmitResponse struct {

    // 商品创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 商品所属市场
    Market   string `json:"market,omitempty"`

}
