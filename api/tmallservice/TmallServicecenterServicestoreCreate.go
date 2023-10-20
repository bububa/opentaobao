package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterservicestorecreate 创建门店
// tmall.servicecenter.servicestore.create
//
// 用于创建门店/网点。多个业务共用
func Tmallservicecenterservicestorecreate(clt *core.SDKClient, req *tmallservice.TmallservicecenterservicestorecreateAPIRequest, session string) (*tmallservice.TmallservicecenterservicestorecreateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterservicestorecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
