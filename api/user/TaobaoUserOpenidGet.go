package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaouseropenidget 查询用户openId
// taobao.user.openid.get
//
// 查询用户openId
func Taobaouseropenidget(clt *core.SDKClient, req *user.TaobaouseropenidgetAPIRequest, session string) (*user.TaobaouseropenidgetAPIResponse, error) {
	var resp user.TaobaouseropenidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
