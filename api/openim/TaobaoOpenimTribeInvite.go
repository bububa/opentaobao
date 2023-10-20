package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribeinvite OPENIM群邀请加入
// taobao.openim.tribe.invite
//
// OPENIM群邀请加入接口
func Taobaoopenimtribeinvite(clt *core.SDKClient, req *openim.TaobaoopenimtribeinviteAPIRequest, session string) (*openim.TaobaoopenimtribeinviteAPIResponse, error) {
	var resp openim.TaobaoopenimtribeinviteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
