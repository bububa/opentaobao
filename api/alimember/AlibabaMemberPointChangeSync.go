package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberPointChangeSync 成长值/积分变更记录同步
// alibaba.member.point.change.sync
//
// 成长值/积分变更记录同步
func AlibabaMemberPointChangeSync(clt *core.SDKClient, req *alimember.AlibabaMemberPointChangeSyncAPIRequest, session string) (*alimember.AlibabaMemberPointChangeSyncAPIResponse, error) {
	var resp alimember.AlibabaMemberPointChangeSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
