package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardpush 推送服务工单信息
// tmall.servicecenter.workcard.push
//
// 服务商家推送工单信息到天猫。
func Tmallservicecenterworkcardpush(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardpushAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardpushAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
