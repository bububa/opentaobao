package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountLoginbytoken 百川TOKEN 登录
// taobao.baichuan.openaccount.loginbytoken
//
// 百川TOKEN 登录
func TaobaoBaichuanOpenaccountLoginbytoken(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLoginbytokenAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountLoginbytokenAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
