package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交客服绩效接口 API请求
taobao.weike.performance.put

提交客服绩效接口
*/
type TaobaoWeikePerformancePutRequest struct {
    model.Params
    // 订单id
    _id   int64
    // 绩效数据封装类
    _perInfoWrapper   *PerformanceInfoWrapper
}

// 初始化TaobaoWeikePerformancePutRequest对象
func NewTaobaoWeikePerformancePutRequest() *TaobaoWeikePerformancePutRequest{
    return &TaobaoWeikePerformancePutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeikePerformancePutRequest) GetApiMethodName() string {
    return "taobao.weike.performance.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeikePerformancePutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 订单id
func (r *TaobaoWeikePerformancePutRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoWeikePerformancePutRequest) GetId() int64 {
    return r._id
}
// PerInfoWrapper Setter
// 绩效数据封装类
func (r *TaobaoWeikePerformancePutRequest) SetPerInfoWrapper(_perInfoWrapper *PerformanceInfoWrapper) error {
    r._perInfoWrapper = _perInfoWrapper
    r.Set("per_info_wrapper", _perInfoWrapper)
    return nil
}

// PerInfoWrapper Getter
func (r TaobaoWeikePerformancePutRequest) GetPerInfoWrapper() *PerformanceInfoWrapper {
    return r._perInfoWrapper
}
