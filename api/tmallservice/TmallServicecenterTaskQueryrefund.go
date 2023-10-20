package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecentertaskqueryrefund 查询任务类工单是否退款
// tmall.servicecenter.task.queryrefund
//
// 查询任务类工单是否退款
func Tmallservicecentertaskqueryrefund(clt *core.SDKClient, req *tmallservice.TmallservicecentertaskqueryrefundAPIRequest, session string) (*tmallservice.TmallservicecentertaskqueryrefundAPIResponse, error) {
	var resp tmallservice.TmallservicecentertaskqueryrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
