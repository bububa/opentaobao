package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmallhkclearanceinfosend 清关信息回调通知
// tmall.hk.clearance.info.send
//
// 清关信息回调通知
func Tmallhkclearanceinfosend(clt *core.SDKClient, req *tmallhk.TmallhkclearanceinfosendAPIRequest, session string) (*tmallhk.TmallhkclearanceinfosendAPIResponse, error) {
	var resp tmallhk.TmallhkclearanceinfosendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
