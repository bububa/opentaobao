package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountresetcodecheck 百川验证找回密码验证码
// taobao.baichuan.openaccount.resetcode.check
//
// 百川验证找回密码验证码
func Taobaobaichuanopenaccountresetcodecheck(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountresetcodecheckAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountresetcodecheckAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountresetcodecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
