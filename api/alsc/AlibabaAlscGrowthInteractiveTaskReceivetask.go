package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscgrowthinteractivetaskreceivetask 领取任务
// alibaba.alsc.growth.interactive.task.receivetask
//
// 领取任务
func Alibabaalscgrowthinteractivetaskreceivetask(clt *core.SDKClient, req *alsc.AlibabaalscgrowthinteractivetaskreceivetaskAPIRequest, session string) (*alsc.AlibabaalscgrowthinteractivetaskreceivetaskAPIResponse, error) {
	var resp alsc.AlibabaalscgrowthinteractivetaskreceivetaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
