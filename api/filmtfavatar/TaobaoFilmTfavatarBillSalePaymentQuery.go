package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// Taobaofilmtfavatarbillsalepaymentquery 获取影院卖品账单--支付账单
// taobao.film.tfavatar.bill.sale.payment.query
//
// 获取影院卖品账单--支付账单
func Taobaofilmtfavatarbillsalepaymentquery(clt *core.SDKClient, req *filmtfavatar.TaobaofilmtfavatarbillsalepaymentqueryAPIRequest, session string) (*filmtfavatar.TaobaofilmtfavatarbillsalepaymentqueryAPIResponse, error) {
	var resp filmtfavatar.TaobaofilmtfavatarbillsalepaymentqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
