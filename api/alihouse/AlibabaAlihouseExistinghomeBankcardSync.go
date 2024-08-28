package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBankcardSync 结算账号同步
// alibaba.alihouse.existinghome.bankcard.sync
//
// 结算账号同步
func AlibabaAlihouseExistinghomeBankcardSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBankcardSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBankcardSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
