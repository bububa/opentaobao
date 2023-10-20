package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamemberexit 退会
// alibaba.member.exit
//
// 商家会员解绑
func Alibabamemberexit(clt *core.SDKClient, req *alimember.AlibabamemberexitAPIRequest, session string) (*alimember.AlibabamemberexitAPIResponse, error) {
	var resp alimember.AlibabamemberexitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
