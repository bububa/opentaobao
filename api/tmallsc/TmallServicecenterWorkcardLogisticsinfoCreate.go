package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardLogisticsinfoCreate 创建服务履约物流单
// tmall.servicecenter.workcard.logisticsinfo.create
//
// 创建服务履约物流单
func TmallServicecenterWorkcardLogisticsinfoCreate(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest, resp *tmallsc.TmallServicecenterWorkcardLogisticsinfoCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
