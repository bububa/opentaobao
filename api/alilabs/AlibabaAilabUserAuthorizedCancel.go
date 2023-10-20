package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabuserauthorizedcancel 取消账号授权
// alibaba.ailab.user.authorized.cancel
//
// 三方用户取消授权给天猫精灵用户
func Alibabaailabuserauthorizedcancel(clt *core.SDKClient, req *alilabs.AlibabaailabuserauthorizedcancelAPIRequest, session string) (*alilabs.AlibabaailabuserauthorizedcancelAPIResponse, error) {
	var resp alilabs.AlibabaailabuserauthorizedcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
