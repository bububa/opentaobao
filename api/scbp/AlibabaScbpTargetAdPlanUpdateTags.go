package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplanupdatetags 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
// alibaba.scbp.target.ad.plan.update.tags
//
// 定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
func Alibabascbptargetadplanupdatetags(clt *core.SDKClient, req *scbp.AlibabascbptargetadplanupdatetagsAPIRequest, session string) (*scbp.AlibabascbptargetadplanupdatetagsAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplanupdatetagsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
