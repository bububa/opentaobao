package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanuserlogin 百川H5登录
// taobao.baichuan.user.login
//
// 百川H5登录
func Taobaobaichuanuserlogin(clt *core.SDKClient, req *baichuan.TaobaobaichuanuserloginAPIRequest, session string) (*baichuan.TaobaobaichuanuserloginAPIResponse, error) {
	var resp baichuan.TaobaobaichuanuserloginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
