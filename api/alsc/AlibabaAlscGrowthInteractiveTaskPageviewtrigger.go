package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscgrowthinteractivetaskpageviewtrigger 浏览打点接口
// alibaba.alsc.growth.interactive.task.pageviewtrigger
//
// 浏览打点接口
func Alibabaalscgrowthinteractivetaskpageviewtrigger(clt *core.SDKClient, req *alsc.AlibabaalscgrowthinteractivetaskpageviewtriggerAPIRequest, session string) (*alsc.AlibabaalscgrowthinteractivetaskpageviewtriggerAPIResponse, error) {
	var resp alsc.AlibabaalscgrowthinteractivetaskpageviewtriggerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
