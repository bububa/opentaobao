package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardLogisticsinfoCreate 创建服务履约物流单
// tmall.servicecenter.workcard.logisticsinfo.create
//
// 创建服务履约物流单
func TmallServicecenterWorkcardLogisticsinfoCreate(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest, session string) (*tmallsc.TmallServicecenterWorkcardLogisticsinfoCreateAPIResponse, error) {
	var resp tmallsc.TmallServicecenterWorkcardLogisticsinfoCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
