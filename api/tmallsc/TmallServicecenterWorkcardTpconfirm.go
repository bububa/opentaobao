package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterworkcardtpconfirm 确认服务完成
// tmall.servicecenter.workcard.tpconfirm
//
// 服务商确认服务完成
func Tmallservicecenterworkcardtpconfirm(clt *core.SDKClient, req *tmallsc.TmallservicecenterworkcardtpconfirmAPIRequest, session string) (*tmallsc.TmallservicecenterworkcardtpconfirmAPIResponse, error) {
	var resp tmallsc.TmallservicecenterworkcardtpconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
