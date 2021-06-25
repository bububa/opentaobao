package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标签人群预估 APIRequest
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

func NewAlibabaScbpAdTargetTagEstimateUvRequest() *AlibabaScbpAdTargetTagEstimateUvRequest{
    return &AlibabaScbpAdTargetTagEstimateUvRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdTargetTagEstimateUvRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.target.tag.estimate.uv"
}

func (r AlibabaScbpAdTargetTagEstimateUvRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdTargetTagEstimateUvRequest) SetTagEstimateOperation(tagEstimateOperation *TagEstimateOperationDto) error {
    r.tagEstimateOperation = tagEstimateOperation
    r.Set("tag_estimate_operation", tagEstimateOperation)
    return nil
}

func (r AlibabaScbpAdTargetTagEstimateUvRequest) GetTagEstimateOperation() *TagEstimateOperationDto {
    return r.tagEstimateOperation
}

func (r *AlibabaScbpAdTargetTagEstimateUvRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdTargetTagEstimateUvRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

