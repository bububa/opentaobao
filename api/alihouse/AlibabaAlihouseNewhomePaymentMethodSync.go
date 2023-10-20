package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomePaymentMethodSync 付款方式上翻
// alibaba.alihouse.newhome.payment.method.sync
//
// 付款方式上翻
func AlibabaAlihouseNewhomePaymentMethodSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomePaymentMethodSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
