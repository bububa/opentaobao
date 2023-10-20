package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaocardexpandcardquery 购物金卡查询
// taobao.card.expandcard.query
//
// 购物金充值信息查询接口，会返回余额等信息。
func Taobaocardexpandcardquery(clt *core.SDKClient, req *promotion.TaobaocardexpandcardqueryAPIRequest, session string) (*promotion.TaobaocardexpandcardqueryAPIResponse, error) {
	var resp promotion.TaobaocardexpandcardqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
