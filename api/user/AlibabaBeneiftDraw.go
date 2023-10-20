package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Alibababeneiftdraw 抽奖接口
// alibaba.beneift.draw
//
// 抽奖接口
func Alibababeneiftdraw(clt *core.SDKClient, req *user.AlibababeneiftdrawAPIRequest, session string) (*user.AlibababeneiftdrawAPIResponse, error) {
	var resp user.AlibababeneiftdrawAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
