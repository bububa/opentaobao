package yunosminiapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosminiapp"
)

// Yunosminiappdatatunnelcall 车载小程序外部服务调用
// yunos.miniapp.datatunnel.call
//
// 对客户提供的api进行统一封装调用。
func Yunosminiappdatatunnelcall(clt *core.SDKClient, req *yunosminiapp.YunosminiappdatatunnelcallAPIRequest, session string) (*yunosminiapp.YunosminiappdatatunnelcallAPIResponse, error) {
	var resp yunosminiapp.YunosminiappdatatunnelcallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
