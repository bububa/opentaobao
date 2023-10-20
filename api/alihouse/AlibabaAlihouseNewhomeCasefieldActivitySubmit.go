package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomecasefieldactivitysubmit 案场活动维护
// alibaba.alihouse.newhome.casefield.activity.submit
//
// 案场活动维护
func Alibabaalihousenewhomecasefieldactivitysubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomecasefieldactivitysubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomecasefieldactivitysubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
