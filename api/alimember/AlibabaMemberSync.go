package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamembersync 会员信息同步
// alibaba.member.sync
//
// 会员信息同步
func Alibabamembersync(clt *core.SDKClient, req *alimember.AlibabamembersyncAPIRequest, session string) (*alimember.AlibabamembersyncAPIResponse, error) {
	var resp alimember.AlibabamembersyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
