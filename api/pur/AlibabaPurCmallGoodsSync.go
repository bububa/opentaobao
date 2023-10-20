package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurCmallGoodsSync 第三方商家接入采购商城-商品同步
// alibaba.pur.cmall.goods.sync
//
// 第三方商家接入采购商城-商品同步
func AlibabaPurCmallGoodsSync(clt *core.SDKClient, req *pur.AlibabaPurCmallGoodsSyncAPIRequest, resp *pur.AlibabaPurCmallGoodsSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
