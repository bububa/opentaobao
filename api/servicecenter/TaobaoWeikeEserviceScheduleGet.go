package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoWeikeEserviceScheduleGet 客服排班信息查询接口
// taobao.weike.eservice.schedule.get
//
// 客服排班信息查询接口
func TaobaoWeikeEserviceScheduleGet(clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceScheduleGetAPIRequest, resp *servicecenter.TaobaoWeikeEserviceScheduleGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
