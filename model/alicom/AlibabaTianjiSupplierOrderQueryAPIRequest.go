package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTianjiSupplierOrderQueryAPIRequest
查询供应商订单 API请求
alibaba.tianji.supplier.order.query

查询供应商订单 */
type AlibabaTianjiSupplierOrderQueryAPIRequest struct {
	model.Params
	// 订单查询入参
	_paramSupplierTopQueryModel *SupplierTopQueryModel
}

// New
