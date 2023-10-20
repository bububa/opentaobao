package pur

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// Alibabapurbasketmerge 合并购物车
// alibaba.pur.basket.merge
//
// 采购商城接入第三方商家合并购物车接口服务
func Alibabapurbasketmerge(clt *core.SDKClient, req *pur.AlibabapurbasketmergeAPIRequest, session string) (*pur.AlibabapurbasketmergeAPIResponse, error) {
	var resp pur.AlibabapurbasketmergeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
