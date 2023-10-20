package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardExpressorderCreate 物流订单创建API
// tmall.servicecenter.workcard.expressorder.create
//
// 天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API
func TmallServicecenterWorkcardExpressorderCreate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardExpressorderCreateAPIRequest, resp *tmallservice.TmallServicecenterWorkcardExpressorderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
