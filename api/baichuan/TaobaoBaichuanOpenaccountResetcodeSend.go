package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountresetcodesend 百川发送找回密码验证码
// taobao.baichuan.openaccount.resetcode.send
//
// 百川发送找回密码验证码
func Taobaobaichuanopenaccountresetcodesend(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountresetcodesendAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountresetcodesendAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountresetcodesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
