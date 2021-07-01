package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

/* TaobaoWeikeEserviceSchedulePut
提交客服排班信息
taobao.weike.eservice.schedule.put

添加、更新、删除排班信息 */
func TaobaoWeikeEserviceSchedulePut(clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceSchedulePutAPIRequest, session string) (*servicecenter.TaobaoWeikeEserviceSchedulePutAPIResponse, error) {
	var resp servicecenter.TaobaoWeikeEserviceSchedulePutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
