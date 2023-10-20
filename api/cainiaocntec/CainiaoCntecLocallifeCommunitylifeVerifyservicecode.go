package cainiaocntec

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// CainiaoCntecLocallifeCommunitylifeVerifyservicecode 验证码验证
// cainiao.cntec.locallife.communitylife.verifyservicecode
//
// 验证码验证
func CainiaoCntecLocallifeCommunitylifeVerifyservicecode(clt *core.SDKClient, req *cainiaocntec.CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest, resp *cainiaocntec.CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
