package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardlogisticsorderupdate 物流工单履约状态更新
// tmall.servicecenter.workcard.logisticsorder.update
//
// 天猫寄送类服务对接外部物流服务商回传物流状态信息
func Tmallservicecenterworkcardlogisticsorderupdate(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardlogisticsorderupdateAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardlogisticsorderupdateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardlogisticsorderupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
