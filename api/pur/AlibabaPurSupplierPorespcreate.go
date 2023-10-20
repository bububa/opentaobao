package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurSupplierPorespcreate po反馈创建
// alibaba.pur.supplier.porespcreate
//
// PO反馈接口
func AlibabaPurSupplierPorespcreate(clt *core.SDKClient, req *pur.AlibabaPurSupplierPorespcreateAPIRequest, resp *pur.AlibabaPurSupplierPorespcreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
