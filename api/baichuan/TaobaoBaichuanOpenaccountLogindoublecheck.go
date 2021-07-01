package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

/* TaobaoBaichuanOpenaccountLogindoublecheck
百川登录二次验证
taobao.baichuan.openaccount.logindoublecheck

百川登录二次验证 */
func TaobaoBaichuanOpenaccountLogindoublecheck(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse, error) {
	var resp baichuan.TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
