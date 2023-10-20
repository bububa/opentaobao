package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountResetcodeCheck 百川验证找回密码验证码
// taobao.baichuan.openaccount.resetcode.check
//
// 百川验证找回密码验证码
func TaobaoBaichuanOpenaccountResetcodeCheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountResetcodeCheckAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountResetcodeCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
