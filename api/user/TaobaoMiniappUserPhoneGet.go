package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaominiappuserphoneget 获取当前授权用户手机号码
// taobao.miniapp.user.phone.get
//
// 在商家应用中，获取当前授权用户手机号码
func Taobaominiappuserphoneget(clt *core.SDKClient, req *user.TaobaominiappuserphonegetAPIRequest, session string) (*user.TaobaominiappuserphonegetAPIResponse, error) {
	var resp user.TaobaominiappuserphonegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
