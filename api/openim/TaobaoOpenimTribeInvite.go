package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeInvite OPENIM群邀请加入
// taobao.openim.tribe.invite
//
// OPENIM群邀请加入接口
func TaobaoOpenimTribeInvite(clt *core.SDKClient, req *openim.TaobaoOpenimTribeInviteAPIRequest, session string) (*openim.TaobaoOpenimTribeInviteAPIResponse, error) {
	var resp openim.TaobaoOpenimTribeInviteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
