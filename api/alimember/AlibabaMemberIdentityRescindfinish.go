package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberIdentityRescindfinish 取消确认
// alibaba.member.identity.rescindfinish
//
// 取消确认
func AlibabaMemberIdentityRescindfinish(clt *core.SDKClient, req *alimember.AlibabaMemberIdentityRescindfinishAPIRequest, session string) (*alimember.AlibabaMemberIdentityRescindfinishAPIResponse, error) {
	var resp alimember.AlibabaMemberIdentityRescindfinishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
