package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimioscertsandboxset 设置开发环境证书
// taobao.openim.ioscert.sandbox.set
//
// 设置开发环境证书
func Taobaoopenimioscertsandboxset(clt *core.SDKClient, req *openim.TaobaoopenimioscertsandboxsetAPIRequest, session string) (*openim.TaobaoopenimioscertsandboxsetAPIResponse, error) {
	var resp openim.TaobaoopenimioscertsandboxsetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
