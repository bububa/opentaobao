package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

/* TaobaoWeikeEserviceScheduleGet
客服排班信息查询接口
taobao.weike.eservice.schedule.get

客服排班信息查询接口 */
func TaobaoWeikeEserviceScheduleGet(clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceScheduleGetAPIRequest, session string) (*servicecenter.TaobaoWeikeEserviceScheduleGetAPIResponse, error) {
	var resp servicecenter.TaobaoWeikeEserviceScheduleGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
