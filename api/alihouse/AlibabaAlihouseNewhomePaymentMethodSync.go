package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomePaymentMethodSync 付款方式上翻
// alibaba.alihouse.newhome.payment.method.sync
//
// 付款方式上翻
func AlibabaAlihouseNewhomePaymentMethodSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomePaymentMethodSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
