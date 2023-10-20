package cainiaocntec

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// CainiaoCntecLocallifeCommunitylifeVerifyservicecode 验证码验证
// cainiao.cntec.locallife.communitylife.verifyservicecode
//
// 验证码验证
func CainiaoCntecLocallifeCommunitylifeVerifyservicecode(clt *core.SDKClient, req *cainiaocntec.CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest, session string) (*cainiaocntec.CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse, error) {
	var resp cainiaocntec.CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
