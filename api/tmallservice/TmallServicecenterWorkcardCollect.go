package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardcollect 工单揽件
// tmall.servicecenter.workcard.collect
//
// 服务工单揽件接口
func Tmallservicecenterworkcardcollect(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardcollectAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardcollectAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardcollectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
