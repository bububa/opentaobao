package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaCeresSupplierPoQuery 采购供应商订单查询接口
// alibaba.ceres.supplier.po.query
//
// 采购供应商订单查询接口
func AlibabaCeresSupplierPoQuery(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaCeresSupplierPoQueryAPIRequest, resp *pur.AlibabaCeresSupplierPoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
