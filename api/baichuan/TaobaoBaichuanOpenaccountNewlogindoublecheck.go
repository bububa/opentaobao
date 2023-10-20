package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountnewlogindoublecheck 百川新登录二次验证
// taobao.baichuan.openaccount.newlogindoublecheck
//
// 百川新登录二次验证
func Taobaobaichuanopenaccountnewlogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountnewlogindoublecheckAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountnewlogindoublecheckAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountnewlogindoublecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
