package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaominiappeleuserphoneget 获取饿了么用户信息
// taobao.miniapp.eleuser.phone.get
//
// 获取饿了么用户信息
func Taobaominiappeleuserphoneget(clt *core.SDKClient, req *user.TaobaominiappeleuserphonegetAPIRequest, session string) (*user.TaobaominiappeleuserphonegetAPIResponse, error) {
	var resp user.TaobaominiappeleuserphonegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
