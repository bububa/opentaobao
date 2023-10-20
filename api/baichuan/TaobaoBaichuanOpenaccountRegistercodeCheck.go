package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountregistercodecheck 百川检查注册验证码
// taobao.baichuan.openaccount.registercode.check
//
// 百川检查注册验证码
func Taobaobaichuanopenaccountregistercodecheck(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountregistercodecheckAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountregistercodecheckAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountregistercodecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
