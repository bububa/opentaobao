package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurSupplierAsncreate asn创建
// alibaba.pur.supplier.asncreate
//
// asn创建
func AlibabaPurSupplierAsncreate(clt *core.SDKClient, req *pur.AlibabaPurSupplierAsncreateAPIRequest, resp *pur.AlibabaPurSupplierAsncreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
