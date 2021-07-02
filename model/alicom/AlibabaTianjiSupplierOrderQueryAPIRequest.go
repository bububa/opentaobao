package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianjiSupplierOrderQueryAPIRequest 查询供应商订单 API请求
// alibaba.tianji.supplier.order.query
//
// 查询供应商订单
type AlibabaTianjiSupplierOrderQueryAPIRequest struct {
	model.Params
	// 订单查询入参
	_paramSupplierTopQueryModel *SupplierTopQueryModel
}

// NewAlibabaTianjiSupplierOrderQueryRequest 初始化AlibabaTianjiSupplierOrderQueryAPIRequest对象
func NewAlibabaTianjiSupplierOrderQueryRequest() *AlibabaTianjiSupplierOrderQueryAPIRequest {
	return &AlibabaTianjiSupplierOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianjiSupplierOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tianji.supplier.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianjiSupplierOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamSupplierTopQueryModel is ParamSupplierTopQueryModel Setter
// 订单查询入参
func (r *AlibabaTianjiSupplierOrderQueryAPIRequest) SetParamSupplierTopQueryModel(_paramSupplierTopQueryModel *SupplierTopQueryModel) error {
	r._paramSupplierTopQueryModel = _paramSupplierTopQueryModel
	r.Set("param_supplier_top_query_model", _paramSupplierTopQueryModel)
	return nil
}

// GetParamSupplierTopQueryModel ParamSupplierTopQueryModel Getter
func (r AlibabaTianjiSupplierOrderQueryAPIRequest) GetParamSupplierTopQueryModel() *SupplierTopQueryModel {
	return r._paramSupplierTopQueryModel
}
