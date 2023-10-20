package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardTpconfirm 确认服务完成
// tmall.servicecenter.workcard.tpconfirm
//
// 服务商确认服务完成
func TmallServicecenterWorkcardTpconfirm(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardTpconfirmAPIRequest, resp *tmallsc.TmallServicecenterWorkcardTpconfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
