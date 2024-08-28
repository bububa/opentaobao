package nazca

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaTokenIssuecertapplyGet 根据token获取出证申请信息
// alibaba.nazca.token.issuecertapply.get
//
// 根据token获取出证申请信息
func AlibabaNazcaTokenIssuecertapplyGet(ctx context.Context, clt *core.SDKClient, req *nazca.AlibabaNazcaTokenIssuecertapplyGetAPIRequest, resp *nazca.AlibabaNazcaTokenIssuecertapplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
