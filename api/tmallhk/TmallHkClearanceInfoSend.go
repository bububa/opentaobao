package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceInfoSend 清关信息回调通知
// tmall.hk.clearance.info.send
//
// 清关信息回调通知
func TmallHkClearanceInfoSend(clt *core.SDKClient, req *tmallhk.TmallHkClearanceInfoSendAPIRequest, resp *tmallhk.TmallHkClearanceInfoSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
