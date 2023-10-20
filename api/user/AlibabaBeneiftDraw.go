package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaBeneiftDraw 抽奖接口
// alibaba.beneift.draw
//
// 抽奖接口
func AlibabaBeneiftDraw(clt *core.SDKClient, req *user.AlibabaBeneiftDrawAPIRequest, resp *user.AlibabaBeneiftDrawAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
