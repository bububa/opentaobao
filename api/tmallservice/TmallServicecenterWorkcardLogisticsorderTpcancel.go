package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardLogisticsorderTpcancel tp更新物流进度信息
// tmall.servicecenter.workcard.logisticsorder.tpcancel
//
// tp更新物流进度信息
func TmallServicecenterWorkcardLogisticsorderTpcancel(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardLogisticsorderTpcancelAPIRequest, resp *tmallservice.TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
