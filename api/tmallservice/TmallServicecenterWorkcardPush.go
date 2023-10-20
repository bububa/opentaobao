package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardPush 推送服务工单信息
// tmall.servicecenter.workcard.push
//
// 服务商家推送工单信息到天猫。
func TmallServicecenterWorkcardPush(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardPushAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardPushAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
