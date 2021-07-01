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
type AlibabaIcbuRfqRecommendAPIRequest struct {
    model.Params
    // 入参数据
    _queryDto   *QueryDTO
}

// 初始化AlibabaIcbuRfqRecommendAPIRequest对象
func NewAlibabaIcbuRfqRecommendRequest() *AlibabaIcbuRfqRecommendAPIRequest{
    return &AlibabaIcbuRfqRecommendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuRfqRecommendAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.recommend"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuRfqRecommendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryDto Setter
// 入参数据
func (r *AlibabaIcbuRfqRecommendAPIRequest) SetQueryDto(_queryDto *QueryDTO) error {
    r._queryDto = _queryDto
    r.Set("query_dto", _queryDto)
    return nil
}

// QueryDto Getter
func (r AlibabaIcbuRfqRecommendAPIRequest) GetQueryDto() *QueryDTO {
    return r._queryDto
}
