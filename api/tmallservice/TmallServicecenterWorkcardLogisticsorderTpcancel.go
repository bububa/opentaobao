package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardLogisticsorderTpcancel tp更新物流进度信息
// tmall.servicecenter.workcard.logisticsorder.tpcancel
//
// tp更新物流进度信息
func TmallServicecenterWorkcardLogisticsorderTpcancel(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
