package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardLogisticsinfoUpdate 物流单信息回传接口
// tmall.servicecenter.workcard.logisticsinfo.update
//
// 物流单信息回传接口
func TmallServicecenterWorkcardLogisticsinfoUpdate(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardLogisticsinfoUpdateAPIRequest, resp *tmallsc.TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
