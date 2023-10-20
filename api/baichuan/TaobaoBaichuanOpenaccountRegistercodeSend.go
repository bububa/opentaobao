package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountregistercodesend 百川发送注册验证码
// taobao.baichuan.openaccount.registercode.send
//
// 百川发送注册验证码
func Taobaobaichuanopenaccountregistercodesend(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountregistercodesendAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountregistercodesendAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountregistercodesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
