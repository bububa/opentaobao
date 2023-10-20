package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberExit 退会
// alibaba.member.exit
//
// 商家会员解绑
func AlibabaMemberExit(clt *core.SDKClient, req *alimember.AlibabaMemberExitAPIRequest, session string) (*alimember.AlibabaMemberExitAPIResponse, error) {
	var resp alimember.AlibabaMemberExitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
