package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaDiafiTokenCheck 天朗token校验API
// alibaba.diafi.token.check
//
// 天朗token校验
func AlibabaDiafiTokenCheck(clt *core.SDKClient, req *security.AlibabaDiafiTokenCheckAPIRequest, resp *security.AlibabaDiafiTokenCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
