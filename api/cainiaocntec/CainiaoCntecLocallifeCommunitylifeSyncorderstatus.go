package cainiaocntec

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// Cainiaocnteclocallifecommunitylifesyncorderstatus 订单状态推送
// cainiao.cntec.locallife.communitylife.syncorderstatus
//
// 订单状态推送
func Cainiaocnteclocallifecommunitylifesyncorderstatus(clt *core.SDKClient, req *cainiaocntec.CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest, session string) (*cainiaocntec.CainiaocnteclocallifecommunitylifesyncorderstatusAPIResponse, error) {
	var resp cainiaocntec.CainiaocnteclocallifecommunitylifesyncorderstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
