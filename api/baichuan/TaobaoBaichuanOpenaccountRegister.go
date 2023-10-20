package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountRegister 百川账号注册
// taobao.baichuan.openaccount.register
//
// 百川账号注册
func TaobaoBaichuanOpenaccountRegister(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountRegisterAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
