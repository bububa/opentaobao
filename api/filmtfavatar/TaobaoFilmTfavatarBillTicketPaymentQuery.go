package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// Taobaofilmtfavatarbillticketpaymentquery 获取影院票务账单-支付订单
// taobao.film.tfavatar.bill.ticket.payment.query
//
// 获取影院票务账单-支付订单
func Taobaofilmtfavatarbillticketpaymentquery(clt *core.SDKClient, req *filmtfavatar.TaobaofilmtfavatarbillticketpaymentqueryAPIRequest, session string) (*filmtfavatar.TaobaofilmtfavatarbillticketpaymentqueryAPIResponse, error) {
	var resp filmtfavatar.TaobaofilmtfavatarbillticketpaymentqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
