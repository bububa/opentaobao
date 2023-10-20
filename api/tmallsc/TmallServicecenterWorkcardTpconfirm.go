package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardTpconfirm 确认服务完成
// tmall.servicecenter.workcard.tpconfirm
//
// 服务商确认服务完成
func TmallServicecenterWorkcardTpconfirm(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardTpconfirmAPIRequest, session string) (*tmallsc.TmallServicecenterWorkcardTpconfirmAPIResponse, error) {
	var resp tmallsc.TmallServicecenterWorkcardTpconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
