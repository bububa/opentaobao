package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountRegistercodeCheck 百川检查注册验证码
// taobao.baichuan.openaccount.registercode.check
//
// 百川检查注册验证码
func TaobaoBaichuanOpenaccountRegistercodeCheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountRegistercodeCheckAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountRegistercodeCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
