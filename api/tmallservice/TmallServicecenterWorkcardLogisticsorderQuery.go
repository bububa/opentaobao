package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardlogisticsorderquery 物流单信息查询
// tmall.servicecenter.workcard.logisticsorder.query
//
// 物流订单信息查询API
func Tmallservicecenterworkcardlogisticsorderquery(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardlogisticsorderqueryAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardlogisticsorderqueryAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardlogisticsorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
