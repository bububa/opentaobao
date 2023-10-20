package pur

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCeresSupplierPoQueryAPIRequest 采购供应商订单查询接口 API请求
// alibaba.ceres.supplier.po.query
//
// 采购供应商订单查询接口
type AlibabaCeresSupplierPoQueryAPIRequest struct {
	model.Params
	// 订单创建日期开始时间
	_startDate string
	// 订单创建日期结束时间
	_endDate string
	// 订单状态
	_status string
}

// NewAlibabaCeresSupplierPoQueryRequest 初始化AlibabaCeresSupplierPoQueryAPIRequest对象
func NewAlibabaCeresSupplierPoQueryRequest() *AlibabaCeresSupplierPoQueryAPIRequest {
	return &AlibabaCeresSupplierPoQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCeresSupplierPoQueryAPIRequest) Reset() {
	r._startDate = ""
	r._endDate = ""
	r._status = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ceres.supplier.po.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 订单创建日期开始时间
func (r *AlibabaCeresSupplierPoQueryAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 订单创建日期结束时间
func (r *AlibabaCeresSupplierPoQueryAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStatus is Status Setter
// 订单状态
func (r *AlibabaCeresSupplierPoQueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetStatus() string {
	return r._status
}

var poolAlibabaCeresSupplierPoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCeresSupplierPoQueryRequest()
	},
}

// GetAlibabaCeresSupplierPoQueryRequest 从 sync.Pool 获取 AlibabaCeresSupplierPoQueryAPIRequest
func GetAlibabaCeresSupplierPoQueryAPIRequest() *AlibabaCeresSupplierPoQueryAPIRequest {
	return poolAlibabaCeresSupplierPoQueryAPIRequest.Get().(*AlibabaCeresSupplierPoQueryAPIRequest)
}

// ReleaseAlibabaCeresSupplierPoQueryAPIRequest 将 AlibabaCeresSupplierPoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaCeresSupplierPoQueryAPIRequest(v *AlibabaCeresSupplierPoQueryAPIRequest) {
	v.Reset()
	poolAlibabaCeresSupplierPoQueryAPIRequest.Put(v)
}
