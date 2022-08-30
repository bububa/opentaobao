package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterFulfiltaskInsuranceAction 供应链保险链路动作
// tmall.servicecenter.fulfiltask.insurance.action
//
// 服务供应链履约链路 保险类业务履约接口
func TmallServicecenterFulfiltaskInsuranceAction(clt *core.SDKClient, req *tmallservice.TmallServicecenterFulfiltaskInsuranceActionAPIRequest, session string) (*tmallservice.TmallServicecenterFulfiltaskInsuranceActionAPIResponse, error) {
	var resp tmallservice.TmallServicecenterFulfiltaskInsuranceActionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
