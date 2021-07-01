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
type AlibabaScbpAdTargetTagEstimateUvAPIRequest struct {
    model.Params
    // 请求参数
    _tagEstimateOperation   *TagEstimateOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdTargetTagEstimateUvAPIRequest对象
func NewAlibabaScbpAdTargetTagEstimateUvRequest() *AlibabaScbpAdTargetTagEstimateUvAPIRequest{
    return &AlibabaScbpAdTargetTagEstimateUvAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.target.tag.estimate.uv"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagEstimateOperation Setter
// 请求参数
func (r *AlibabaScbpAdTargetTagEstimateUvAPIRequest) SetTagEstimateOperation(_tagEstimateOperation *TagEstimateOperationDTO) error {
    r._tagEstimateOperation = _tagEstimateOperation
    r.Set("tag_estimate_operation", _tagEstimateOperation)
    return nil
}

// TagEstimateOperation Getter
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetTagEstimateOperation() *TagEstimateOperationDTO {
    return r._tagEstimateOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagEstimateUvAPIRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
