package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeactivitysubscriptionbind 销售活动绑定认购商品
// alibaba.alihouse.newhome.activity.subscription.bind
//
// 销售活动绑定认购商品
func Alibabaalihousenewhomeactivitysubscriptionbind(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeactivitysubscriptionbindAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeactivitysubscriptionbindAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeactivitysubscriptionbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
