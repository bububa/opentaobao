package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardexpressordercreate 物流订单创建API
// tmall.servicecenter.workcard.expressorder.create
//
// 天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API
func Tmallservicecenterworkcardexpressordercreate(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardexpressordercreateAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardexpressordercreateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardexpressordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
