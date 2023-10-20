package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabUserAuthorizedCancel 取消账号授权
// alibaba.ailab.user.authorized.cancel
//
// 三方用户取消授权给天猫精灵用户
func AlibabaAilabUserAuthorizedCancel(clt *core.SDKClient, req *alilabs.AlibabaAilabUserAuthorizedCancelAPIRequest, session string) (*alilabs.AlibabaAilabUserAuthorizedCancelAPIResponse, error) {
	var resp alilabs.AlibabaAilabUserAuthorizedCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
