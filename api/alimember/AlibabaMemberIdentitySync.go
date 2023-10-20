package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberIdentitySync 会员身份信息同步
// alibaba.member.identity.sync
//
// 会员身份信息同步
func AlibabaMemberIdentitySync(clt *core.SDKClient, req *alimember.AlibabaMemberIdentitySyncAPIRequest, resp *alimember.AlibabaMemberIdentitySyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
