package nazca

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaTokenChangeauthapplyGet 根据token获取变更认证申请信息
// alibaba.nazca.token.changeauthapply.get
//
// 根据token获取变更认证申请信息
func AlibabaNazcaTokenChangeauthapplyGet(ctx context.Context, clt *core.SDKClient, req *nazca.AlibabaNazcaTokenChangeauthapplyGetAPIRequest, resp *nazca.AlibabaNazcaTokenChangeauthapplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
