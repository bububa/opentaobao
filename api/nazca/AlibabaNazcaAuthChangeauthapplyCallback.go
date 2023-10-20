package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// Alibabanazcaauthchangeauthapplycallback 变更认证回调
// alibaba.nazca.auth.changeauthapply.callback
//
// 变更认证回调
func Alibabanazcaauthchangeauthapplycallback(clt *core.SDKClient, req *nazca.AlibabanazcaauthchangeauthapplycallbackAPIRequest, session string) (*nazca.AlibabanazcaauthchangeauthapplycallbackAPIResponse, error) {
	var resp nazca.AlibabanazcaauthchangeauthapplycallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
