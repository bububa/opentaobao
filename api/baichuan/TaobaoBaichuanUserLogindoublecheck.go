package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanUserLogindoublecheck 百川H5登录二次验证
// taobao.baichuan.user.logindoublecheck
//
// 百川H5登录二次验证
func TaobaoBaichuanUserLogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLogindoublecheckAPIRequest, resp *baichuan.TaobaoBaichuanUserLogindoublecheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
