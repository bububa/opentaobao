package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

/* TaobaoBaichuanUserLogin
百川H5登录
taobao.baichuan.user.login

百川H5登录 */
func TaobaoBaichuanUserLogin(clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLoginAPIRequest, session string) (*baichuan.TaobaoBaichuanUserLoginAPIResponse, error) {
	var resp baichuan.TaobaoBaichuanUserLoginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
