package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardquerybyseller 工单查询接口（面向商家）
// tmall.servicecenter.workcard.querybyseller
//
// 查询工单
func Tmallservicecenterworkcardquerybyseller(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardquerybysellerAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardquerybysellerAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardquerybysellerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
