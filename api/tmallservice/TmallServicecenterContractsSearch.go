package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecentercontractssearch 获取合同类的服务工单信息
// tmall.servicecenter.contracts.search
//
// 获取合同类的服务工单信息
func Tmallservicecentercontractssearch(clt *core.SDKClient, req *tmallservice.TmallservicecentercontractssearchAPIRequest, session string) (*tmallservice.TmallservicecentercontractssearchAPIResponse, error) {
	var resp tmallservice.TmallservicecentercontractssearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
