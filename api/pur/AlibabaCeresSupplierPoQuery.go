package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaCeresSupplierPoQuery 采购供应商订单查询接口
// alibaba.ceres.supplier.po.query
//
// 采购供应商订单查询接口
func AlibabaCeresSupplierPoQuery(clt *core.SDKClient, req *pur.AlibabaCeresSupplierPoQueryAPIRequest, resp *pur.AlibabaCeresSupplierPoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
