package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// Taobaouseravatarget 淘宝用户头像查询
// taobao.user.avatar.get
//
// 根据混淆nick查询用户头像
func Taobaouseravatarget(clt *core.SDKClient, req *tbuser.TaobaouseravatargetAPIRequest, session string) (*tbuser.TaobaouseravatargetAPIResponse, error) {
	var resp tbuser.TaobaouseravatargetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
