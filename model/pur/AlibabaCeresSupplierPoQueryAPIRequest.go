package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCeresSupplierPoQueryAPIRequest
采购供应商订单查询接口 API请求
alibaba.ceres.supplier.po.query

采购供应商订单查询接口 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ceres.supplier.po.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StartDate Setter
// 订单创建日期开始时间
func (r *AlibabaCeresSupplierPoQueryAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 订单创建日期结束时间
func (r *AlibabaCeresSupplierPoQueryAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is Status Setter
// 订单状态
func (r *AlibabaCeresSupplierPoQueryAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaCeresSupplierPoQueryAPIRequest) GetStatus() string {
	return r._status
}
