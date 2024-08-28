package nazca

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaTokenAuthapplyGet 根据token获取认证申请信息
// alibaba.nazca.token.authapply.get
//
// 根据token获取认证申请信息
func AlibabaNazcaTokenAuthapplyGet(ctx context.Context, clt *core.SDKClient, req *nazca.AlibabaNazcaTokenAuthapplyGetAPIRequest, resp *nazca.AlibabaNazcaTokenAuthapplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
