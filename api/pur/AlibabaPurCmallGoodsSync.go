package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapurcmallgoodssync 第三方商家接入采购商城-商品同步
// alibaba.pur.cmall.goods.sync
//
// 第三方商家接入采购商城-商品同步
func Alibabapurcmallgoodssync(clt *core.SDKClient, req *pur.AlibabapurcmallgoodssyncAPIRequest, session string) (*pur.AlibabapurcmallgoodssyncAPIResponse, error) {
	var resp pur.AlibabapurcmallgoodssyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
