package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeactivitysave 新增或者更新销售活动
// alibaba.alihouse.newhome.activity.save
//
// 新增或者更新销售活动
func Alibabaalihousenewhomeactivitysave(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeactivitysaveAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeactivitysaveAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeactivitysaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
