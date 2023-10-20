package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMiniappUserPhoneGet 获取当前授权用户手机号码
// taobao.miniapp.user.phone.get
//
// 在商家应用中，获取当前授权用户手机号码
func TaobaoMiniappUserPhoneGet(clt *core.SDKClient, req *user.TaobaoMiniappUserPhoneGetAPIRequest, resp *user.TaobaoMiniappUserPhoneGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
