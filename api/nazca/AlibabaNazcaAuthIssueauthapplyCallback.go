package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// Alibabanazcaauthissueauthapplycallback 出证申请回调
// alibaba.nazca.auth.issueauthapply.callback
//
// 出证申请回调
func Alibabanazcaauthissueauthapplycallback(clt *core.SDKClient, req *nazca.AlibabanazcaauthissueauthapplycallbackAPIRequest, session string) (*nazca.AlibabanazcaauthissueauthapplycallbackAPIResponse, error) {
	var resp nazca.AlibabanazcaauthissueauthapplycallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
