package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymSteamShadowuserGet 获取影子标识
// alibaba.jym.steam.shadowuser.get
//
// 交易猫Steam类目获取影子ID
func AlibabaJymSteamShadowuserGet(clt *core.SDKClient, req *jym.AlibabaJymSteamShadowuserGetAPIRequest, resp *jym.AlibabaJymSteamShadowuserGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
