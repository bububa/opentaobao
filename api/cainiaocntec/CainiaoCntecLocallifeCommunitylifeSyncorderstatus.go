package cainiaocntec

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// CainiaoCntecLocallifeCommunitylifeSyncorderstatus 订单状态推送
// cainiao.cntec.locallife.communitylife.syncorderstatus
//
// 订单状态推送
func CainiaoCntecLocallifeCommunitylifeSyncorderstatus(clt *core.SDKClient, req *cainiaocntec.CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest, session string) (*cainiaocntec.CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponse, error) {
	var resp cainiaocntec.CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
