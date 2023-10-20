package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsFulfillQaOrderCreateAPIRequest 创单接口 API请求
// alibaba.tcls.fulfill.qa.order.create
//
// 根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据）
type AlibabaTclsFulfillQaOrderCreateAPIRequest struct {
	model.Params
	// 原始履约单号
	_fulfillOrderId string
	// 目标ip
	_targetIp string
	// 执行人姓名
	_creator string
	// 执行人工号
	_jobNo string
}

// NewAlibabaTclsFulfillQaOrderCreateRequest 初始化AlibabaTclsFulfillQaOrderCreateAPIRequest对象
func NewAlibabaTclsFulfillQaOrderCreateRequest() *AlibabaTclsFulfillQaOrderCreateAPIRequest {
	return &AlibabaTclsFulfillQaOrderCreateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) Reset() {
	r._fulfillOrderId = ""
	r._targetIp = ""
	r._creator = ""
	r._jobNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.fulfill.qa.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillOrderId is FulfillOrderId Setter
// 原始履约单号
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) SetFulfillOrderId(_fulfillOrderId string) error {
	r._fulfillOrderId = _fulfillOrderId
	r.Set("fulfill_order_id", _fulfillOrderId)
	return nil
}

// GetFulfillOrderId FulfillOrderId Getter
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetFulfillOrderId() string {
	return r._fulfillOrderId
}

// SetTargetIp is TargetIp Setter
// 目标ip
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) SetTargetIp(_targetIp string) error {
	r._targetIp = _targetIp
	r.Set("target_ip", _targetIp)
	return nil
}

// GetTargetIp TargetIp Getter
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetTargetIp() string {
	return r._targetIp
}

// SetCreator is Creator Setter
// 执行人姓名
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) SetCreator(_creator string) error {
	r._creator = _creator
	r.Set("creator", _creator)
	return nil
}

// GetCreator Creator Getter
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetCreator() string {
	return r._creator
}

// SetJobNo is JobNo Setter
// 执行人工号
func (r *AlibabaTclsFulfillQaOrderCreateAPIRequest) SetJobNo(_jobNo string) error {
	r._jobNo = _jobNo
	r.Set("job_no", _jobNo)
	return nil
}

// GetJobNo JobNo Getter
func (r AlibabaTclsFulfillQaOrderCreateAPIRequest) GetJobNo() string {
	return r._jobNo
}

var poolAlibabaTclsFulfillQaOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsFulfillQaOrderCreateRequest()
	},
}

// GetAlibabaTclsFulfillQaOrderCreateRequest 从 sync.Pool 获取 AlibabaTclsFulfillQaOrderCreateAPIRequest
func GetAlibabaTclsFulfillQaOrderCreateAPIRequest() *AlibabaTclsFulfillQaOrderCreateAPIRequest {
	return poolAlibabaTclsFulfillQaOrderCreateAPIRequest.Get().(*AlibabaTclsFulfillQaOrderCreateAPIRequest)
}

// ReleaseAlibabaTclsFulfillQaOrderCreateAPIRequest 将 AlibabaTclsFulfillQaOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsFulfillQaOrderCreateAPIRequest(v *AlibabaTclsFulfillQaOrderCreateAPIRequest) {
	v.Reset()
	poolAlibabaTclsFulfillQaOrderCreateAPIRequest.Put(v)
}
