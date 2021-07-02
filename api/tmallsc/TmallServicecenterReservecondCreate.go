package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterReservecondCreate 创建主动预约开通条件
// tmall.servicecenter.reservecond.create
//
// 1、设置主动预约开通条件
func TmallServicecenterReservecondCreate(clt *core.SDKClient, req *tmallsc.TmallServicecenterReservecondCreateAPIRequest, session string) (*tmallsc.TmallServicecenterReservecondCreateAPIResponse, error) {
	var resp tmallsc.TmallServicecenterReservecondCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
