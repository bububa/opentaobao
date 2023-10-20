package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountloginbytoken 百川TOKEN 登录
// taobao.baichuan.openaccount.loginbytoken
//
// 百川TOKEN 登录
func Taobaobaichuanopenaccountloginbytoken(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountloginbytokenAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountloginbytokenAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountloginbytokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
