package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// Taobaofilmtfavatarbillsalerefundqueryvii 获取影院卖品账单--退款账单-V2版本(正逆分离)
// taobao.film.tfavatar.bill.sale.refund.query.vii
//
// 获取影院卖品账单--退款账单-V2版本(正逆分离)
func Taobaofilmtfavatarbillsalerefundqueryvii(clt *core.SDKClient, req *filmtfavatar.TaobaofilmtfavatarbillsalerefundqueryviiAPIRequest, session string) (*filmtfavatar.TaobaofilmtfavatarbillsalerefundqueryviiAPIResponse, error) {
	var resp filmtfavatar.TaobaofilmtfavatarbillsalerefundqueryviiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
