package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// Taobaofilmtfavatarbillsalerefundquery 获取影院卖品账单--退款账单
// taobao.film.tfavatar.bill.sale.refund.query
//
// 获取影院卖品账单--退款账单
func Taobaofilmtfavatarbillsalerefundquery(clt *core.SDKClient, req *filmtfavatar.TaobaofilmtfavatarbillsalerefundqueryAPIRequest, session string) (*filmtfavatar.TaobaofilmtfavatarbillsalerefundqueryAPIResponse, error) {
	var resp filmtfavatar.TaobaofilmtfavatarbillsalerefundqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
