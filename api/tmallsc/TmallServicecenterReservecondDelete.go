package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterreserveconddelete 删除主动预约开通条件
// tmall.servicecenter.reservecond.delete
//
// 删除主动预约开通条件
func Tmallservicecenterreserveconddelete(clt *core.SDKClient, req *tmallsc.TmallservicecenterreserveconddeleteAPIRequest, session string) (*tmallsc.TmallservicecenterreserveconddeleteAPIResponse, error) {
	var resp tmallsc.TmallservicecenterreserveconddeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
