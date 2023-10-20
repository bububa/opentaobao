package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeactivitydepositunbind 销售活动解绑预存金商品
// alibaba.alihouse.newhome.activity.deposit.unbind
//
// 销售活动解绑预存金商品
func Alibabaalihousenewhomeactivitydepositunbind(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeactivitydepositunbindAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeactivitydepositunbindAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeactivitydepositunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
