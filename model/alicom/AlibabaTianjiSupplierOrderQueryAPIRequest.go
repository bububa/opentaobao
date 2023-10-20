package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianjisupplierorderqueryAPIRequest 查询供应商订单 API请求
// alibaba.tianji.supplier.order.query
//
// 查询供应商订单
type AlibabatianjisupplierorderqueryAPIRequest struct {
	model.Params
	// 订单查询入参
	_paramSupplierTopQueryModel *SupplierTopQueryModel
}

// NewAlibabatianjisupplierorderqueryRequest 初始化AlibabatianjisupplierorderqueryAPIRequest对象
func NewAlibabatianjisupplierorderqueryRequest() *AlibabatianjisupplierorderqueryAPIRequest {
	return &AlibabatianjisupplierorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianjisupplierorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tianji.supplier.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianjisupplierorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianjisupplierorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSupplierTopQueryModel is ParamSupplierTopQueryModel Setter
// 订单查询入参
func (r *AlibabatianjisupplierorderqueryAPIRequest) SetParamSupplierTopQueryModel(_paramSupplierTopQueryModel *SupplierTopQueryModel) error {
	r._paramSupplierTopQueryModel = _paramSupplierTopQueryModel
	r.Set("param_supplier_top_query_model", _paramSupplierTopQueryModel)
	return nil
}

// GetParamSupplierTopQueryModel ParamSupplierTopQueryModel Getter
func (r AlibabatianjisupplierorderqueryAPIRequest) GetParamSupplierTopQueryModel() *SupplierTopQueryModel {
	return r._paramSupplierTopQueryModel
}
