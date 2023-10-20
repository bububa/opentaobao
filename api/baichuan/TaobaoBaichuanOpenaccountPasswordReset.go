package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountpasswordreset 百川找回密码
// taobao.baichuan.openaccount.password.reset
//
// 百川找回密码
func Taobaobaichuanopenaccountpasswordreset(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountpasswordresetAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountpasswordresetAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountpasswordresetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
