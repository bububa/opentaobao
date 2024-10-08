package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeSubAccountBind 子账号入驻
// alibaba.alihouse.existinghome.sub.account.bind
//
// 子账号入驻
func AlibabaAlihouseExistinghomeSubAccountBind(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeSubAccountBindAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeSubAccountBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
