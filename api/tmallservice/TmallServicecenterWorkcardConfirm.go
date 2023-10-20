package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardconfirm 服务商确认服务完成
// tmall.servicecenter.workcard.confirm
//
// 提供给外部合作服务商，用于通知天猫，告知寄修服务厂内操作全部完成
func Tmallservicecenterworkcardconfirm(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardconfirmAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardconfirmAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
