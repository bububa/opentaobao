package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeactivitycustomersave 销售活动批量保存定向用户
// alibaba.alihouse.newhome.activity.customer.save
//
// 销售活动批量保存定向用户
func Alibabaalihousenewhomeactivitycustomersave(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeactivitycustomersaveAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeactivitycustomersaveAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeactivitycustomersaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
