package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurSupplierPorespcreate po反馈创建
// alibaba.pur.supplier.porespcreate
//
// PO反馈接口
func AlibabaPurSupplierPorespcreate(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurSupplierPorespcreateAPIRequest, resp *pur.AlibabaPurSupplierPorespcreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
