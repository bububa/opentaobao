package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscgrowthinteractivetaskreceivetaskprize 任务领奖
// alibaba.alsc.growth.interactive.task.receivetaskprize
//
// 任务领奖
func Alibabaalscgrowthinteractivetaskreceivetaskprize(clt *core.SDKClient, req *alsc.AlibabaalscgrowthinteractivetaskreceivetaskprizeAPIRequest, session string) (*alsc.AlibabaalscgrowthinteractivetaskreceivetaskprizeAPIResponse, error) {
	var resp alsc.AlibabaalscgrowthinteractivetaskreceivetaskprizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
