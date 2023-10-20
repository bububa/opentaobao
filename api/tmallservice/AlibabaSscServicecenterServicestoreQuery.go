package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabasscservicecenterservicestorequery 根据天猫id查询门店信息
// alibaba.ssc.servicecenter.servicestore.query
//
// 根据天猫id查询门店信息
func Alibabasscservicecenterservicestorequery(clt *core.SDKClient, req *tmallservice.AlibabasscservicecenterservicestorequeryAPIRequest, session string) (*tmallservice.AlibabasscservicecenterservicestorequeryAPIResponse, error) {
	var resp tmallservice.AlibabasscservicecenterservicestorequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
