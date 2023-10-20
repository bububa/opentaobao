package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurProductSync 同步产品
// alibaba.pur.product.sync
//
// 同步产品
func AlibabaPurProductSync(clt *core.SDKClient, req *pur.AlibabaPurProductSyncAPIRequest, resp *pur.AlibabaPurProductSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
