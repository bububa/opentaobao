package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
rfq推荐 APIRequest
alibaba.icbu.rfq.recommend

rfq推荐
*/
type AlibabaIcbuRfqRecommendRequest struct {
    model.Params

    // 入参数据
    queryDto   *QueryDto 

}

func NewAlibabaIcbuRfqRecommendRequest() *AlibabaIcbuRfqRecommendRequest{
    return &AlibabaIcbuRfqRecommendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuRfqRecommendRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.recommend"
}

func (r AlibabaIcbuRfqRecommendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuRfqRecommendRequest) SetQueryDto(queryDto *QueryDto) error {
    r.queryDto = queryDto
    r.Set("query_dto", queryDto)
    return nil
}

func (r AlibabaIcbuRfqRecommendRequest) GetQueryDto() *QueryDto {
    return r.queryDto
}

