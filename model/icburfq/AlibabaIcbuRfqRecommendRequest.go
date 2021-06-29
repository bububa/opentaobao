package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rfq推荐 API请求
alibaba.icbu.rfq.recommend

rfq推荐
*/
type AlibabaIcbuRfqRecommendRequest struct {
    model.Params
    // 入参数据
    _queryDto   *QueryDto
}

// 初始化AlibabaIcbuRfqRecommendRequest对象
func NewAlibabaIcbuRfqRecommendRequest() *AlibabaIcbuRfqRecommendRequest{
    return &AlibabaIcbuRfqRecommendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqRecommendRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.recommend"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqRecommendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryDto Setter
// 入参数据
func (r *AlibabaIcbuRfqRecommendRequest) SetQueryDto(_queryDto *QueryDto) error {
    r._queryDto = _queryDto
    r.Set("query_dto", _queryDto)
    return nil
}

// QueryDto Getter
func (r AlibabaIcbuRfqRecommendRequest) GetQueryDto() *QueryDto {
    return r._queryDto
}
