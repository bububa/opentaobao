package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaceressupplierpoqueryAPIRequest 采购供应商订单查询接口 API请求
// alibaba.ceres.supplier.po.query
//
// 采购供应商订单查询接口
type AlibabaceressupplierpoqueryAPIRequest struct {
	model.Params
	// 订单创建日期开始时间
	_startDate string
	// 订单创建日期结束时间
	_endDate string
	// 订单状态
	_status string
}

// NewAlibabaceressupplierpoqueryRequest 初始化AlibabaceressupplierpoqueryAPIRequest对象
func NewAlibabaceressupplierpoqueryRequest() *AlibabaceressupplierpoqueryAPIRequest {
	return &AlibabaceressupplierpoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaceressupplierpoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ceres.supplier.po.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaceressupplierpoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaceressupplierpoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 订单创建日期开始时间
func (r *AlibabaceressupplierpoqueryAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaceressupplierpoqueryAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 订单创建日期结束时间
func (r *AlibabaceressupplierpoqueryAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaceressupplierpoqueryAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStatus is Status Setter
// 订单状态
func (r *AlibabaceressupplierpoqueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaceressupplierpoqueryAPIRequest) GetStatus() string {
	return r._status
}
