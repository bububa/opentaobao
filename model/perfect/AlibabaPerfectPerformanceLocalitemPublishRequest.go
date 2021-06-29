package perfect

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城购定制化发品 APIRequest
alibaba.perfect.performance.localitem.publish

同城购业务定制化发品接口，同城购业务线专用
*/
type AlibabaPerfectPerformanceLocalitemPublishRequest struct {
    model.Params

    // 请求参数
    paramPerfectPerformanceItemPublishReq   *PerfectPerformanceItemPublishReq 

}

func NewAlibabaPerfectPerformanceLocalitemPublishRequest() *AlibabaPerfectPerformanceLocalitemPublishRequest{
    return &AlibabaPerfectPerformanceLocalitemPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPerfectPerformanceLocalitemPublishRequest) GetApiMethodName() string {
    return "alibaba.perfect.performance.localitem.publish"
}

func (r AlibabaPerfectPerformanceLocalitemPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPerfectPerformanceLocalitemPublishRequest) SetParamPerfectPerformanceItemPublishReq(paramPerfectPerformanceItemPublishReq *PerfectPerformanceItemPublishReq) error {
    r.paramPerfectPerformanceItemPublishReq = paramPerfectPerformanceItemPublishReq
    r.Set("param_perfect_performance_item_publish_req", paramPerfectPerformanceItemPublishReq)
    return nil
}

func (r AlibabaPerfectPerformanceLocalitemPublishRequest) GetParamPerfectPerformanceItemPublishReq() *PerfectPerformanceItemPublishReq {
    return r.paramPerfectPerformanceItemPublishReq
}

