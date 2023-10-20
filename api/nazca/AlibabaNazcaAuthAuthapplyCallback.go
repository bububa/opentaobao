package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// Alibabanazcaauthauthapplycallback 认证的统一回调接口
// alibaba.nazca.auth.authapply.callback
//
// 认证的统一回调接口
func Alibabanazcaauthauthapplycallback(clt *core.SDKClient, req *nazca.AlibabanazcaauthauthapplycallbackAPIRequest, session string) (*nazca.AlibabanazcaauthauthapplycallbackAPIResponse, error) {
	var resp nazca.AlibabanazcaauthauthapplycallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
