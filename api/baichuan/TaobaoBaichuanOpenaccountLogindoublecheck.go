package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountLogindoublecheck 百川登录二次验证
// taobao.baichuan.openaccount.logindoublecheck
//
// 百川登录二次验证
func TaobaoBaichuanOpenaccountLogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
