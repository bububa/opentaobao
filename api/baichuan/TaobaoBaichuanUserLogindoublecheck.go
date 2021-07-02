package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanUserLogindoublecheck 百川H5登录二次验证
// taobao.baichuan.user.logindoublecheck
//
// 百川H5登录二次验证
func TaobaoBaichuanUserLogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLogindoublecheckAPIRequest, session string) (*baichuan.TaobaoBaichuanUserLogindoublecheckAPIResponse, error) {
	var resp baichuan.TaobaoBaichuanUserLogindoublecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
