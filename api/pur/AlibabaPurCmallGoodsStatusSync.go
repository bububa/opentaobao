package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurCmallGoodsStatusSync 第三方商城接入采购商城-商品状态同步
// alibaba.pur.cmall.goods.status.sync
//
// 第三方商城接入采购商城-商品状态同步
func AlibabaPurCmallGoodsStatusSync(clt *core.SDKClient, req *pur.AlibabaPurCmallGoodsStatusSyncAPIRequest, session string) (*pur.AlibabaPurCmallGoodsStatusSyncAPIResponse, error) {
	var resp pur.AlibabaPurCmallGoodsStatusSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
