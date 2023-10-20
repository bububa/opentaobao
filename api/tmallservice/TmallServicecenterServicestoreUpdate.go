package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterservicestoreupdate 修改门店信息
// tmall.servicecenter.servicestore.update
//
// 用于修改门店/网点信息。多个业务共用
func Tmallservicecenterservicestoreupdate(clt *core.SDKClient, req *tmallservice.TmallservicecenterservicestoreupdateAPIRequest, session string) (*tmallservice.TmallservicecenterservicestoreupdateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterservicestoreupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
