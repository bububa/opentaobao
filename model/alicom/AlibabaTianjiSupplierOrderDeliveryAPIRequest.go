package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTianjiSupplierOrderDeliveryAPIRequest
天机供应商发货 API请求
alibaba.tianji.supplier.order.delivery

天机供应商发货 */
type AlibabaTianjiSupplierOrderDeliveryAPIRequest struct {
	model.Params
	// 物流信息
	_paramDistributionOrderLogisticsModel *DistributionOrderLogisticsModel
}

// New
