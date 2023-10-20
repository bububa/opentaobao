package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimIoscertSandboxSet 设置开发环境证书
// taobao.openim.ioscert.sandbox.set
//
// 设置开发环境证书
func TaobaoOpenimIoscertSandboxSet(clt *core.SDKClient, req *openim.TaobaoOpenimIoscertSandboxSetAPIRequest, resp *openim.TaobaoOpenimIoscertSandboxSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
