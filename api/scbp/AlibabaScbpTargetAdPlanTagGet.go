package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplantagget 定向推广-获取计划的定向溢价数据
// alibaba.scbp.target.ad.plan.tag.get
//
// 定向推广-获取计划的定向溢价数据
func Alibabascbptargetadplantagget(clt *core.SDKClient, req *scbp.AlibabascbptargetadplantaggetAPIRequest, session string) (*scbp.AlibabascbptargetadplantaggetAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplantaggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
