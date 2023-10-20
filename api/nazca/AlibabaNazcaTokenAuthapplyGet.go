package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaTokenAuthapplyGet 根据token获取认证申请信息
// alibaba.nazca.token.authapply.get
//
// 根据token获取认证申请信息
func AlibabaNazcaTokenAuthapplyGet(clt *core.SDKClient, req *nazca.AlibabaNazcaTokenAuthapplyGetAPIRequest, resp *nazca.AlibabaNazcaTokenAuthapplyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
