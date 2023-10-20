package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanuserlogindoublecheck 百川H5登录二次验证
// taobao.baichuan.user.logindoublecheck
//
// 百川H5登录二次验证
func Taobaobaichuanuserlogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaobaichuanuserlogindoublecheckAPIRequest, session string) (*baichuan.TaobaobaichuanuserlogindoublecheckAPIResponse, error) {
	var resp baichuan.TaobaobaichuanuserlogindoublecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
