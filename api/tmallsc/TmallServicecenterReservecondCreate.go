package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterreservecondcreate 创建主动预约开通条件
// tmall.servicecenter.reservecond.create
//
// 1、设置主动预约开通条件
func Tmallservicecenterreservecondcreate(clt *core.SDKClient, req *tmallsc.TmallservicecenterreservecondcreateAPIRequest, session string) (*tmallsc.TmallservicecenterreservecondcreateAPIResponse, error) {
	var resp tmallsc.TmallservicecenterreservecondcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
