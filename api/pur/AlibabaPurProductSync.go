package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapurproductsync 同步产品
// alibaba.pur.product.sync
//
// 同步产品
func Alibabapurproductsync(clt *core.SDKClient, req *pur.AlibabapurproductsyncAPIRequest, session string) (*pur.AlibabapurproductsyncAPIResponse, error) {
	var resp pur.AlibabapurproductsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
