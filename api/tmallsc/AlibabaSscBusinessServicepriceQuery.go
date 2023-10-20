package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaSscBusinessServicepriceQuery 苹果查询服务供给报价
// alibaba.ssc.business.serviceprice.query
//
// 苹果查询服务供给报价
func AlibabaSscBusinessServicepriceQuery(clt *core.SDKClient, req *tmallsc.AlibabaSscBusinessServicepriceQueryAPIRequest, resp *tmallsc.AlibabaSscBusinessServicepriceQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
