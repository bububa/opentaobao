package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplanoperation 定向推广-计划开启/暂停/删除
// alibaba.scbp.target.ad.plan.operation
//
// 定向推广-计划开启/暂停/删除
func Alibabascbptargetadplanoperation(clt *core.SDKClient, req *scbp.AlibabascbptargetadplanoperationAPIRequest, session string) (*scbp.AlibabascbptargetadplanoperationAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplanoperationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
