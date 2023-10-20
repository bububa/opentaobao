package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurBasketMerge 合并购物车
// alibaba.pur.basket.merge
//
// 采购商城接入第三方商家合并购物车接口服务
func AlibabaPurBasketMerge(clt *core.SDKClient, req *pur.AlibabaPurBasketMergeAPIRequest, session string) (*pur.AlibabaPurBasketMergeAPIResponse, error) {
	var resp pur.AlibabaPurBasketMergeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
