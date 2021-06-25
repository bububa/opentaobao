package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交客服绩效接口 APIRequest
taobao.weike.performance.put

提交客服绩效接口
*/
type TaobaoWeikePerformancePutRequest struct {
    model.Params

    // 订单id
    id   int64 

    // 绩效数据封装类
    perInfoWrapper   *PerformanceInfoWrapper 

}

func NewTaobaoWeikePerformancePutRequest() *TaobaoWeikePerformancePutRequest{
    return &TaobaoWeikePerformancePutRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWeikePerformancePutRequest) GetApiMethodName() string {
    return "taobao.weike.performance.put"
}

func (r TaobaoWeikePerformancePutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWeikePerformancePutRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoWeikePerformancePutRequest) GetId() int64 {
    return r.id
}

func (r *TaobaoWeikePerformancePutRequest) SetPerInfoWrapper(perInfoWrapper *PerformanceInfoWrapper) error {
    r.perInfoWrapper = perInfoWrapper
    r.Set("per_info_wrapper", perInfoWrapper)
    return nil
}

func (r TaobaoWeikePerformancePutRequest) GetPerInfoWrapper() *PerformanceInfoWrapper {
    return r.perInfoWrapper
}

