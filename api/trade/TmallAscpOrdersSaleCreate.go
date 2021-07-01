package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

/* TmallAscpOrdersSaleCreate
ASCP渠道中心销售单创建接口
tmall.ascp.orders.sale.create

ASCP渠道中心销售单创建接口 */
func TmallAscpOrdersSaleCreate(clt *core.SDKClient, req *trade.TmallAscpOrdersSaleCreateAPIRequest, session string) (*trade.TmallAscpOrdersSaleCreateAPIResponse, error) {
	var resp trade.TmallAscpOrdersSaleCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
