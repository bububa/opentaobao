package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务工单拉取 API请求
tmall.servicecenter.task.get

接口供服务供应商通过交易主订单查询其未拉取的任务类工单
*/
type TmallServicecenterTaskGetRequest struct {
    model.Params
    // Taobao主交易订单ID
    _parentBizOrderId   int64
}

// 初始化TmallServicecenterTaskGetRequest对象
func NewTmallServicecenterTaskGetRequest() *TmallServicecenterTaskGetRequest{
    return &TmallServicecenterTaskGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterTaskGetRequest) GetApiMethodName() string {
    return "tmall.servicecenter.task.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterTaskGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParentBizOrderId Setter
// Taobao主交易订单ID
func (r *TmallServicecenterTaskGetRequest) SetParentBizOrderId(_parentBizOrderId int64) error {
    r._parentBizOrderId = _parentBizOrderId
    r.Set("parent_biz_order_id", _parentBizOrderId)
    return nil
}

// ParentBizOrderId Getter
func (r TmallServicecenterTaskGetRequest) GetParentBizOrderId() int64 {
    return r._parentBizOrderId
}
