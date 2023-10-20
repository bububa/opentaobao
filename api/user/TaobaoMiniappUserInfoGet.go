package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaominiappuserInfoget 用户开放信息获取
// taobao.miniapp.userInfo.get
//
// 获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
func TaobaominiappuserInfoget(clt *core.SDKClient, req *user.TaobaominiappuserInfogetAPIRequest, session string) (*user.TaobaominiappuserInfogetAPIResponse, error) {
	var resp user.TaobaominiappuserInfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
