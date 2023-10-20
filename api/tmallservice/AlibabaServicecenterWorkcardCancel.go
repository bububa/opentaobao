package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterWorkcardCancel 服务平台工单取消接口
// alibaba.servicecenter.workcard.cancel
//
// 取消服务工单
func AlibabaServicecenterWorkcardCancel(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterWorkcardCancelAPIRequest, resp *tmallservice.AlibabaServicecenterWorkcardCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
