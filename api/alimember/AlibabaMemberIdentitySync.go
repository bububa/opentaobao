package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamemberidentitysync 会员身份信息同步
// alibaba.member.identity.sync
//
// 会员身份信息同步
func Alibabamemberidentitysync(clt *core.SDKClient, req *alimember.AlibabamemberidentitysyncAPIRequest, session string) (*alimember.AlibabamemberidentitysyncAPIResponse, error) {
	var resp alimember.AlibabamemberidentitysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
