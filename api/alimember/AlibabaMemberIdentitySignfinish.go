package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberIdentitySignfinish 签约确认
// alibaba.member.identity.signfinish
//
// 签约确认
func AlibabaMemberIdentitySignfinish(clt *core.SDKClient, req *alimember.AlibabaMemberIdentitySignfinishAPIRequest, resp *alimember.AlibabaMemberIdentitySignfinishAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
