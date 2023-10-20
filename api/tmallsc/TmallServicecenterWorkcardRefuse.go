package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardRefuse 买家拒收
// tmall.servicecenter.workcard.refuse
//
// 买家拒收通知接口
func TmallServicecenterWorkcardRefuse(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardRefuseAPIRequest, resp *tmallsc.TmallServicecenterWorkcardRefuseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
