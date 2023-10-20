package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianjisupplierorderresultAPIRequest 供应商处理订单接口（订购成功/失败、发货） API请求
// alibaba.tianji.supplier.order.result
//
// 供应商处理订单接口（订购成功/失败、发货）
type AlibabatianjisupplierorderresultAPIRequest struct {
	model.Params
	// 供应商处理订单结果反馈参数
	_supplierOrderResultModel *SupplierOrderResultModel
}

// NewAlibabatianjisupplierorderresultRequest 初始化AlibabatianjisupplierorderresultAPIRequest对象
func NewAlibabatianjisupplierorderresultRequest() *AlibabatianjisupplierorderresultAPIRequest {
	return &AlibabatianjisupplierorderresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianjisupplierorderresultAPIRequest) GetApiMethodName() string {
	return "alibaba.tianji.supplier.order.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianjisupplierorderresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianjisupplierorderresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierOrderResultModel is SupplierOrderResultModel Setter
// 供应商处理订单结果反馈参数
func (r *AlibabatianjisupplierorderresultAPIRequest) SetSupplierOrderResultModel(_supplierOrderResultModel *SupplierOrderResultModel) error {
	r._supplierOrderResultModel = _supplierOrderResultModel
	r.Set("supplier_order_result_model", _supplierOrderResultModel)
	return nil
}

// GetSupplierOrderResultModel SupplierOrderResultModel Getter
func (r AlibabatianjisupplierorderresultAPIRequest) GetSupplierOrderResultModel() *SupplierOrderResultModel {
	return r._supplierOrderResultModel
}
