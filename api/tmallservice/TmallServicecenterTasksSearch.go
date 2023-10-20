package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecentertaskssearch 查询任务类工单信息
// tmall.servicecenter.tasks.search
//
// 查询任务类工单信息
func Tmallservicecentertaskssearch(clt *core.SDKClient, req *tmallservice.TmallservicecentertaskssearchAPIRequest, session string) (*tmallservice.TmallservicecentertaskssearchAPIResponse, error) {
	var resp tmallservice.TmallservicecentertaskssearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
