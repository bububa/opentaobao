package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabasscbusinessservicepricequery 苹果查询服务供给报价
// alibaba.ssc.business.serviceprice.query
//
// 苹果查询服务供给报价
func Alibabasscbusinessservicepricequery(clt *core.SDKClient, req *tmallsc.AlibabasscbusinessservicepricequeryAPIRequest, session string) (*tmallsc.AlibabasscbusinessservicepricequeryAPIResponse, error) {
	var resp tmallsc.AlibabasscbusinessservicepricequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
