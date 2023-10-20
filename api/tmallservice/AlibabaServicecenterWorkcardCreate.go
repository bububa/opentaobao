package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterWorkcardCreate 服务平台工单创建接口
// alibaba.servicecenter.workcard.create
//
// 创建服务平台工单
func AlibabaServicecenterWorkcardCreate(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterWorkcardCreateAPIRequest, resp *tmallservice.AlibabaServicecenterWorkcardCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
