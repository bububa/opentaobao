package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardexpressorderconsign 物流订单呼叫运力
// tmall.servicecenter.workcard.expressorder.consign
//
// 天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
func Tmallservicecenterworkcardexpressorderconsign(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardexpressorderconsignAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardexpressorderconsignAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardexpressorderconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
