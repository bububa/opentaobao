package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceInfoSend 清关信息回调通知
// tmall.hk.clearance.info.send
//
// 清关信息回调通知
func TmallHkClearanceInfoSend(clt *core.SDKClient, req *tmallhk.TmallHkClearanceInfoSendAPIRequest, session string) (*tmallhk.TmallHkClearanceInfoSendAPIResponse, error) {
	var resp tmallhk.TmallHkClearanceInfoSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
