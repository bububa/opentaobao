package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标签人群预估 API请求
alibaba.scbp.ad.target.tag.estimate.uv

标签人群预估
*/
type AlibabaScbpAdTargetTagEstimateUvRequest struct {
    model.Params
    // 请求参数
    tagEstimateOperation   *TagEstimateOperationDto
    // 用户信息
    topContext   *TopContextDto
}

// 初始化AlibabaScbpAdTargetTagEstimateUvRequest对象
func NewAlibabaScbpAdTargetTagEstimateUvRequest() *AlibabaScbpAdTargetTagEstimateUvRequest{
    return &AlibabaScbpAdTargetTagEstimateUvRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagEstimateUvRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.target.tag.estimate.uv"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagEstimateUvRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagEstimateOperation Setter
// 请求参数
func (r *AlibabaScbpAdTargetTagEstimateUvRequest) SetTagEstimateOperation(tagEstimateOperation *TagEstimateOperationDto) error {
    r.tagEstimateOperation = tagEstimateOperation
    r.Set("tag_estimate_operation", tagEstimateOperation)
    return nil
}

// TagEstimateOperation Getter
func (r AlibabaScbpAdTargetTagEstimateUvRequest) GetTagEstimateOperation() *TagEstimateOperationDto {
    return r.tagEstimateOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagEstimateUvRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdTargetTagEstimateUvRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
