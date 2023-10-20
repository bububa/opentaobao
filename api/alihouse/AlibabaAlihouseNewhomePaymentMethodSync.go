package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomePaymentMethodSync 付款方式上翻
// alibaba.alihouse.newhome.payment.method.sync
//
// 付款方式上翻
func AlibabaAlihouseNewhomePaymentMethodSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomePaymentMethodSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomePaymentMethodSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
