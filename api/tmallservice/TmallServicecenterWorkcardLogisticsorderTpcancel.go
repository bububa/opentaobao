package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardlogisticsordertpcancel tp更新物流进度信息
// tmall.servicecenter.workcard.logisticsorder.tpcancel
//
// tp更新物流进度信息
func Tmallservicecenterworkcardlogisticsordertpcancel(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardlogisticsordertpcancelAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardlogisticsordertpcancelAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardlogisticsordertpcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
