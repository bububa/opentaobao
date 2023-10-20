package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsfulfillqaordercreateAPIRequest 创单接口 API请求
// alibaba.tcls.fulfill.qa.order.create
//
// 根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据）
type AlibabatclsfulfillqaordercreateAPIRequest struct {
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

// NewAlibabatclsfulfillqaordercreateRequest 初始化AlibabatclsfulfillqaordercreateAPIRequest对象
func NewAlibabatclsfulfillqaordercreateRequest() *AlibabatclsfulfillqaordercreateAPIRequest {
	return &AlibabatclsfulfillqaordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsfulfillqaordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.fulfill.qa.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsfulfillqaordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsfulfillqaordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillOrderId is FulfillOrderId Setter
// 原始履约单号
func (r *AlibabatclsfulfillqaordercreateAPIRequest) SetFulfillOrderId(_fulfillOrderId string) error {
	r._fulfillOrderId = _fulfillOrderId
	r.Set("fulfill_order_id", _fulfillOrderId)
	return nil
}

// GetFulfillOrderId FulfillOrderId Getter
func (r AlibabatclsfulfillqaordercreateAPIRequest) GetFulfillOrderId() string {
	return r._fulfillOrderId
}

// SetTargetIp is TargetIp Setter
// 目标ip
func (r *AlibabatclsfulfillqaordercreateAPIRequest) SetTargetIp(_targetIp string) error {
	r._targetIp = _targetIp
	r.Set("target_ip", _targetIp)
	return nil
}

// GetTargetIp TargetIp Getter
func (r AlibabatclsfulfillqaordercreateAPIRequest) GetTargetIp() string {
	return r._targetIp
}

// SetCreator is Creator Setter
// 执行人姓名
func (r *AlibabatclsfulfillqaordercreateAPIRequest) SetCreator(_creator string) error {
	r._creator = _creator
	r.Set("creator", _creator)
	return nil
}

// GetCreator Creator Getter
func (r AlibabatclsfulfillqaordercreateAPIRequest) GetCreator() string {
	return r._creator
}

// SetJobNo is JobNo Setter
// 执行人工号
func (r *AlibabatclsfulfillqaordercreateAPIRequest) SetJobNo(_jobNo string) error {
	r._jobNo = _jobNo
	r.Set("job_no", _jobNo)
	return nil
}

// GetJobNo JobNo Getter
func (r AlibabatclsfulfillqaordercreateAPIRequest) GetJobNo() string {
	return r._jobNo
}
