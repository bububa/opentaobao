package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountregister 百川账号注册
// taobao.baichuan.openaccount.register
//
// 百川账号注册
func Taobaobaichuanopenaccountregister(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountregisterAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountregisterAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
