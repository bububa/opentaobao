package yunosminiapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosminiapp"
)

// YunosMiniappActivityCall 调用活动接口
// yunos.miniapp.activity.call
//
// 用于小程序调用活动接口
func YunosMiniappActivityCall(clt *core.SDKClient, req *yunosminiapp.YunosMiniappActivityCallAPIRequest, resp *yunosminiapp.YunosMiniappActivityCallAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
