package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼商品发布 API返回值 
alibaba.idle.isv.item.publish

服务商ISV闲鱼商品发布
*/
type AlibabaIdleIsvItemPublishAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvItemPublishAPIResponseModel
}

// 服务商闲鱼商品发布 成功返回结果
type AlibabaIdleIsvItemPublishAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_item_publish_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *IdleResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
