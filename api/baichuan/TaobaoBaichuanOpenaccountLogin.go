package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountlogin 百川用户名密码登录
// taobao.baichuan.openaccount.login
//
// 百川用户名密码登录
func Taobaobaichuanopenaccountlogin(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountloginAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountloginAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountloginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
