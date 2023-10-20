package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaTianjiSupplierOrderQuery 查询供应商订单
// alibaba.tianji.supplier.order.query
//
// 查询供应商订单
func AlibabaTianjiSupplierOrderQuery(clt *core.SDKClient, req *alicom.AlibabaTianjiSupplierOrderQueryAPIRequest, resp *alicom.AlibabaTianjiSupplierOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
