package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountNewlogindoublecheck 百川新登录二次验证
// taobao.baichuan.openaccount.newlogindoublecheck
//
// 百川新登录二次验证
func TaobaoBaichuanOpenaccountNewlogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountNewlogindoublecheckAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountNewlogindoublecheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
