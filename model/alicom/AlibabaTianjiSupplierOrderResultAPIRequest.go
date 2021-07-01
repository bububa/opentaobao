package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTianjiSupplierOrderResultAPIRequest
供应商处理订单接口（订购成功/失败、发货） API请求
alibaba.tianji.supplier.order.result

供应商处理订单接口（订购成功/失败、发货） */
type AlibabaTianjiSupplierOrderResultAPIRequest struct {
	model.Params
	// 供应商处理订单结果反馈参数
	_supplierOrderResultModel *SupplierOrderResultModel
}

// New
