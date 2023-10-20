package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanopenaccountlogindoublecheck 百川登录二次验证
// taobao.baichuan.openaccount.logindoublecheck
//
// 百川登录二次验证
func Taobaobaichuanopenaccountlogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaobaichuanopenaccountlogindoublecheckAPIRequest, session string) (*baichuan.TaobaobaichuanopenaccountlogindoublecheckAPIResponse, error) {
	var resp baichuan.TaobaobaichuanopenaccountlogindoublecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
