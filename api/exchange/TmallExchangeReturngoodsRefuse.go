package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

/* TmallExchangeReturngoodsRefuse
卖家拒绝确认收货
tmall.exchange.returngoods.refuse

卖家拒绝买家换货申请 */
func TmallExchangeReturngoodsRefuse(clt *core.SDKClient, req *exchange.TmallExchangeReturngoodsRefuseAPIRequest, session string) (*exchange.TmallExchangeReturngoodsRefuseAPIResponse, error) {
	var resp exchange.TmallExchangeReturngoodsRefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
