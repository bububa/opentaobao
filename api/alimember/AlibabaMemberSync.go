package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberSync 会员信息同步
// alibaba.member.sync
//
// 会员信息同步
func AlibabaMemberSync(clt *core.SDKClient, req *alimember.AlibabaMemberSyncAPIRequest, session string) (*alimember.AlibabaMemberSyncAPIResponse, error) {
	var resp alimember.AlibabaMemberSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
