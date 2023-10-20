package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymsteamfulfillmentupdate 交易猫Steam类目发履约态变更
// alibaba.jym.steam.fulfillment.update
//
// 交易猫Steam类目发履约态变更
func Alibabajymsteamfulfillmentupdate(clt *core.SDKClient, req *jym.AlibabajymsteamfulfillmentupdateAPIRequest, session string) (*jym.AlibabajymsteamfulfillmentupdateAPIResponse, error) {
	var resp jym.AlibabajymsteamfulfillmentupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
