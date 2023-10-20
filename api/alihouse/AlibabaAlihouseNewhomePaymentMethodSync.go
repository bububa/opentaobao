package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomepaymentmethodsync 付款方式上翻
// alibaba.alihouse.newhome.payment.method.sync
//
// 付款方式上翻
func Alibabaalihousenewhomepaymentmethodsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomepaymentmethodsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomepaymentmethodsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomepaymentmethodsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
