package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaCeresSupplierPoQuerydetail 采购供应商订单明细查询接口
// alibaba.ceres.supplier.po.querydetail
//
// 采购供应商订单明细查询接口
func AlibabaCeresSupplierPoQuerydetail(clt *core.SDKClient, req *pur.AlibabaCeresSupplierPoQuerydetailAPIRequest, resp *pur.AlibabaCeresSupplierPoQuerydetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
