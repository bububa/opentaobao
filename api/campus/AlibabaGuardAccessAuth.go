package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaGuardAccessAuth 鉴权
// alibaba.guard.access.auth
//
// 刷卡鉴权
func AlibabaGuardAccessAuth(clt *core.SDKClient, req *campus.AlibabaGuardAccessAuthAPIRequest, resp *campus.AlibabaGuardAccessAuthAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
