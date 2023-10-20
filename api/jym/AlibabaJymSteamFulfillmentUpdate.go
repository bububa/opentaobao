package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymSteamFulfillmentUpdate 交易猫Steam类目发履约态变更
// alibaba.jym.steam.fulfillment.update
//
// 交易猫Steam类目发履约态变更
func AlibabaJymSteamFulfillmentUpdate(clt *core.SDKClient, req *jym.AlibabaJymSteamFulfillmentUpdateAPIRequest, session string) (*jym.AlibabaJymSteamFulfillmentUpdateAPIResponse, error) {
	var resp jym.AlibabaJymSteamFulfillmentUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
