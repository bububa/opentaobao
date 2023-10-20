package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanUserLoginbytoken 百川手淘信任登录
// taobao.baichuan.user.loginbytoken
//
// 百川手淘信任登录
func TaobaoBaichuanUserLoginbytoken(clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLoginbytokenAPIRequest, resp *baichuan.TaobaoBaichuanUserLoginbytokenAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
