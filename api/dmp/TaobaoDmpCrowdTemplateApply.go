package dmp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dmp"
)

// TaobaoDmpCrowdTemplateApply 人群模版采纳并生成人群API
// taobao.dmp.crowd.template.apply
//
// 人群模版采纳并生成人群API
func TaobaoDmpCrowdTemplateApply(clt *core.SDKClient, req *dmp.TaobaoDmpCrowdTemplateApplyAPIRequest, session string) (*dmp.TaobaoDmpCrowdTemplateApplyAPIResponse, error) {
	var resp dmp.TaobaoDmpCrowdTemplateApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
