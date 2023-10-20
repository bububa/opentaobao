package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscRightTokenCheck 实物奖品凭证校验
// alibaba.alsc.right.token.check
//
// 实物奖品凭证校验
func AlibabaAlscRightTokenCheck(clt *core.SDKClient, req *alsc.AlibabaAlscRightTokenCheckAPIRequest, resp *alsc.AlibabaAlscRightTokenCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
