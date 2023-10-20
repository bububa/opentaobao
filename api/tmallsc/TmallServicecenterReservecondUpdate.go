package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterreservecondupdate 主动预约条件更新
// tmall.servicecenter.reservecond.update
//
// 1、设置主动预约开通条件
func Tmallservicecenterreservecondupdate(clt *core.SDKClient, req *tmallsc.TmallservicecenterreservecondupdateAPIRequest, session string) (*tmallsc.TmallservicecenterreservecondupdateAPIResponse, error) {
	var resp tmallsc.TmallservicecenterreservecondupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
