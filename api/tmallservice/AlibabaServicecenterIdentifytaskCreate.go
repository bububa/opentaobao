package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterIdentifytaskCreate 创建核销单
// alibaba.servicecenter.identifytask.create
//
// 创建核销单
func AlibabaServicecenterIdentifytaskCreate(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterIdentifytaskCreateAPIRequest, resp *tmallservice.AlibabaServicecenterIdentifytaskCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
