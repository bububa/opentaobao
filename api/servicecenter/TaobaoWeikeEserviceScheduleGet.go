package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaoweikeeservicescheduleget 客服排班信息查询接口
// taobao.weike.eservice.schedule.get
//
// 客服排班信息查询接口
func Taobaoweikeeservicescheduleget(clt *core.SDKClient, req *servicecenter.TaobaoweikeeserviceschedulegetAPIRequest, session string) (*servicecenter.TaobaoweikeeserviceschedulegetAPIResponse, error) {
	var resp servicecenter.TaobaoweikeeserviceschedulegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
