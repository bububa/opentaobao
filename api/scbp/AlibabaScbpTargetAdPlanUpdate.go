package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplanupdate 定向推广-更新推广计划的基础信息
// alibaba.scbp.target.ad.plan.update
//
// 定向推广-更新推广计划的基础信息
func Alibabascbptargetadplanupdate(clt *core.SDKClient, req *scbp.AlibabascbptargetadplanupdateAPIRequest, session string) (*scbp.AlibabascbptargetadplanupdateAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplanupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
