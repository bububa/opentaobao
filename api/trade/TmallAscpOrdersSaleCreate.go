package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TmallAscpOrdersSaleCreate ASCP渠道中心销售单创建接口
// tmall.ascp.orders.sale.create
//
// ASCP渠道中心销售单创建接口
func TmallAscpOrdersSaleCreate(clt *core.SDKClient, req *trade.TmallAscpOrdersSaleCreateAPIRequest, resp *trade.TmallAscpOrdersSaleCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
