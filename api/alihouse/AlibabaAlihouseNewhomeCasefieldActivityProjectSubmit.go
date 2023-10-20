package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomecasefieldactivityprojectsubmit 案场活动楼盘维护
// alibaba.alihouse.newhome.casefield.activity.project.submit
//
// 案场活动楼盘维护
func Alibabaalihousenewhomecasefieldactivityprojectsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
