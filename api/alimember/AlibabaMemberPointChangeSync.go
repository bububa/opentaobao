package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamemberpointchangesync 成长值/积分变更记录同步
// alibaba.member.point.change.sync
//
// 成长值/积分变更记录同步
func Alibabamemberpointchangesync(clt *core.SDKClient, req *alimember.AlibabamemberpointchangesyncAPIRequest, session string) (*alimember.AlibabamemberpointchangesyncAPIResponse, error) {
	var resp alimember.AlibabamemberpointchangesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
