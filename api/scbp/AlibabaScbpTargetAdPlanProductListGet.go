package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplanproductlistget 定向推广-获取推广计划产品列表
// alibaba.scbp.target.ad.plan.product.list.get
//
// 定向推广-获取推广计划产品列表
func Alibabascbptargetadplanproductlistget(clt *core.SDKClient, req *scbp.AlibabascbptargetadplanproductlistgetAPIRequest, session string) (*scbp.AlibabascbptargetadplanproductlistgetAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplanproductlistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
