package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// Taobaofilmtfavatarbillsalepaymentqueryvii 获取影院卖品账单--支付账单-V2版本(正逆分离)
// taobao.film.tfavatar.bill.sale.payment.query.vii
//
// 获取影院卖品账单--支付账单-V2版本(正逆分离)
func Taobaofilmtfavatarbillsalepaymentqueryvii(clt *core.SDKClient, req *filmtfavatar.TaobaofilmtfavatarbillsalepaymentqueryviiAPIRequest, session string) (*filmtfavatar.TaobaofilmtfavatarbillsalepaymentqueryviiAPIResponse, error) {
	var resp filmtfavatar.TaobaofilmtfavatarbillsalepaymentqueryviiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
