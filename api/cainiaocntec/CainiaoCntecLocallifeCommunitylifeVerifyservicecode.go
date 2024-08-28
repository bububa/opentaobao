package cainiaocntec

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// CainiaoCntecLocallifeCommunitylifeVerifyservicecode 验证码验证
// cainiao.cntec.locallife.communitylife.verifyservicecode
//
// 验证码验证
func CainiaoCntecLocallifeCommunitylifeVerifyservicecode(ctx context.Context, clt *core.SDKClient, req *cainiaocntec.CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest, resp *cainiaocntec.CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
