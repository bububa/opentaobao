package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplanadd 定向推广-新建计划
// alibaba.scbp.target.ad.plan.add
//
// 定向推广-新建单条计划
func Alibabascbptargetadplanadd(clt *core.SDKClient, req *scbp.AlibabascbptargetadplanaddAPIRequest, session string) (*scbp.AlibabascbptargetadplanaddAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplanaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
