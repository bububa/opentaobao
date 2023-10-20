package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplanforbiddenwordmodify 定向推广-新增或删除屏蔽词
// alibaba.scbp.target.ad.plan.forbidden.word.modify
//
// 定向推广-新增或删除屏蔽词
func Alibabascbptargetadplanforbiddenwordmodify(clt *core.SDKClient, req *scbp.AlibabascbptargetadplanforbiddenwordmodifyAPIRequest, session string) (*scbp.AlibabascbptargetadplanforbiddenwordmodifyAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplanforbiddenwordmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
