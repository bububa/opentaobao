package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeactivitydepositbind 销售活动绑定预存金商品id
// alibaba.alihouse.newhome.activity.deposit.bind
//
// 销售活动绑定预存金商品id
func Alibabaalihousenewhomeactivitydepositbind(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeactivitydepositbindAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeactivitydepositbindAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeactivitydepositbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
