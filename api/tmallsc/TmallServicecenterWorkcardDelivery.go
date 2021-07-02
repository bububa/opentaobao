package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardDelivery 开始配送工单
// tmall.servicecenter.workcard.delivery
//
// 服务商调用该接口通知天猫服务平台服务商工人已开始配送工单
func TmallServicecenterWorkcardDelivery(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardDeliveryAPIRequest, session string) (*tmallsc.TmallServicecenterWorkcardDeliveryAPIResponse, error) {
	var resp tmallsc.TmallServicecenterWorkcardDeliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
