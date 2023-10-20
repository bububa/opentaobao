package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountLogin 百川用户名密码登录
// taobao.baichuan.openaccount.login
//
// 百川用户名密码登录
func TaobaoBaichuanOpenaccountLogin(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLoginAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountLoginAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
