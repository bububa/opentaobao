package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplanupdateproducts 定向推广 按照id操作推广计划的产品，包括新增，删除和更新
// alibaba.scbp.target.ad.plan.update.products
//
// 定向推广 按照id操作推广计划的产品，包括新增，删除和更新
func Alibabascbptargetadplanupdateproducts(clt *core.SDKClient, req *scbp.AlibabascbptargetadplanupdateproductsAPIRequest, session string) (*scbp.AlibabascbptargetadplanupdateproductsAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplanupdateproductsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
