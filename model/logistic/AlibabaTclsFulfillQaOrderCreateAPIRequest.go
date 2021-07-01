package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创单接口 API请求
alibaba.tcls.fulfill.qa.order.create

根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据）
*/
type AlibabaTclsFulfillQaOrderCreateAPIRequest struct {
    model.Params
    // 原始履约单号
    _fulfillOrderId   string
    // 目标ip
    _targetIp   string
    // 执行人姓名
    _creator   string
    // 执行人工号
    _jobNo   string
}

// 初始化AlibabaTclsFulfillQaOrderCreateAPIRequest对象
func NewAlibabaTclsFulfillQaOrderCreateRequest() *AlibabaTclsFulfillQaOrderCreateAPIRequest{
    return &AlibabaTclsFulfillQaOrderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.tcls.fulfill.qa.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FulfillOrderId Setter
// 原始履约单号
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) SetFulfillOrderId(_fulfillOrderId string) error {
    r._fulfillOrderId = _fulfillOrderId
    r.Set("fulfill_order_id", _fulfillOrderId)
    return nil
}

// FulfillOrderId Getter
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetFulfillOrderId() string {
    return r._fulfillOrderId
}
// TargetIp Setter
// 目标ip
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) SetTargetIp(_targetIp string) error {
    r._targetIp = _targetIp
    r.Set("target_ip", _targetIp)
    return nil
}

// TargetIp Getter
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetTargetIp() string {
    return r._targetIp
}
// Creator Setter
// 执行人姓名
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) SetCreator(_creator string) error {
    r._creator = _creator
    r.Set("creator", _creator)
    return nil
}

// Creator Getter
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetCreator() string {
    return r._creator
}
// JobNo Setter
// 执行人工号
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) SetJobNo(_jobNo string) error {
    r._jobNo = _jobNo
    r.Set("job_no", _jobNo)
    return nil
}

// JobNo Getter
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetJobNo() string {
    return r._jobNo
}
