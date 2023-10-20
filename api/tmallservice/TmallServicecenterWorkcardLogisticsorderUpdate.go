package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardLogisticsorderUpdate 物流工单履约状态更新
// tmall.servicecenter.workcard.logisticsorder.update
//
// 天猫寄送类服务对接外部物流服务商回传物流状态信息
func TmallServicecenterWorkcardLogisticsorderUpdate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest, resp *tmallservice.TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
