package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterworkcardlogisticsinfoupdate 物流单信息回传接口
// tmall.servicecenter.workcard.logisticsinfo.update
//
// 物流单信息回传接口
func Tmallservicecenterworkcardlogisticsinfoupdate(clt *core.SDKClient, req *tmallsc.TmallservicecenterworkcardlogisticsinfoupdateAPIRequest, session string) (*tmallsc.TmallservicecenterworkcardlogisticsinfoupdateAPIResponse, error) {
	var resp tmallsc.TmallservicecenterworkcardlogisticsinfoupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
