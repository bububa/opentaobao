package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplanlist 定向推广-查询定向推广计划列表并返回计划基础信息
// alibaba.scbp.target.ad.plan.list
//
// 定向推广-查询定向推广计划列表并返回计划基础信息
func Alibabascbptargetadplanlist(clt *core.SDKClient, req *scbp.AlibabascbptargetadplanlistAPIRequest, session string) (*scbp.AlibabascbptargetadplanlistAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplanlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
