package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterworkcardlogisticsinfocreate 创建服务履约物流单
// tmall.servicecenter.workcard.logisticsinfo.create
//
// 创建服务履约物流单
func Tmallservicecenterworkcardlogisticsinfocreate(clt *core.SDKClient, req *tmallsc.TmallservicecenterworkcardlogisticsinfocreateAPIRequest, session string) (*tmallsc.TmallservicecenterworkcardlogisticsinfocreateAPIResponse, error) {
	var resp tmallsc.TmallservicecenterworkcardlogisticsinfocreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
