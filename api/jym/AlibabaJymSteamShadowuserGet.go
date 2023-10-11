package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymSteamShadowuserGet 获取影子标识
// alibaba.jym.steam.shadowuser.get
//
// 交易猫Steam类目获取影子ID
func AlibabaJymSteamShadowuserGet(clt *core.SDKClient, req *jym.AlibabaJymSteamShadowuserGetAPIRequest, session string) (*jym.AlibabaJymSteamShadowuserGetAPIResponse, error) {
	var resp jym.AlibabaJymSteamShadowuserGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
