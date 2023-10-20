package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecentermsftasksupdate 喵师傅工人任务批量完成接口
// tmall.servicecenter.msf.tasks.update
//
// 喵师傅工人任务批量完成接口
func Tmallservicecentermsftasksupdate(clt *core.SDKClient, req *tmallservice.TmallservicecentermsftasksupdateAPIRequest, session string) (*tmallservice.TmallservicecentermsftasksupdateAPIResponse, error) {
	var resp tmallservice.TmallservicecentermsftasksupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
