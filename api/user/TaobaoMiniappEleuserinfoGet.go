package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaominiappeleuserinfoget 获取饿了么用户信息
// taobao.miniapp.eleuserinfo.get
//
// 获取饿了么用户信息
func Taobaominiappeleuserinfoget(clt *core.SDKClient, req *user.TaobaominiappeleuserinfogetAPIRequest, session string) (*user.TaobaominiappeleuserinfogetAPIResponse, error) {
	var resp user.TaobaominiappeleuserinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
