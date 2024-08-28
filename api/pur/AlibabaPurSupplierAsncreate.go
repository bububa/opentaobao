package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurSupplierAsncreate asn创建
// alibaba.pur.supplier.asncreate
//
// asn创建
func AlibabaPurSupplierAsncreate(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurSupplierAsncreateAPIRequest, resp *pur.AlibabaPurSupplierAsncreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
