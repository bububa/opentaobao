package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecentertaskget 服务工单拉取
// tmall.servicecenter.task.get
//
// 接口供服务供应商通过交易主订单查询其未拉取的任务类工单
func Tmallservicecentertaskget(clt *core.SDKClient, req *tmallservice.TmallservicecentertaskgetAPIRequest, session string) (*tmallservice.TmallservicecentertaskgetAPIResponse, error) {
	var resp tmallservice.TmallservicecentertaskgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
