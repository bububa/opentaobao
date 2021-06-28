package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品发布 APIResponse
alibaba.item.publish.submit

新商品发布，提交商品发布信息
*/
type AlibabaItemPublishSubmitAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_item_publish_submit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品创建时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"