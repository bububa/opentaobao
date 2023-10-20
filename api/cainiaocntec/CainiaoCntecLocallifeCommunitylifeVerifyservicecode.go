package cainiaocntec

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// Cainiaocnteclocallifecommunitylifeverifyservicecode 验证码验证
// cainiao.cntec.locallife.communitylife.verifyservicecode
//
// 验证码验证
func Cainiaocnteclocallifecommunitylifeverifyservicecode(clt *core.SDKClient, req *cainiaocntec.CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest, session string) (*cainiaocntec.CainiaocnteclocallifecommunitylifeverifyservicecodeAPIResponse, error) {
	var resp cainiaocntec.CainiaocnteclocallifecommunitylifeverifyservicecodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
