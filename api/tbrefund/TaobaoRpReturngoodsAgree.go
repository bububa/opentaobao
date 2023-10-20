package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorpreturngoodsagree 卖家同意退货
// taobao.rp.returngoods.agree
//
// 卖家同意退货，支持淘宝和天猫的订单。
func Taobaorpreturngoodsagree(clt *core.SDKClient, req *tbrefund.TaobaorpreturngoodsagreeAPIRequest, session string) (*tbrefund.TaobaorpreturngoodsagreeAPIResponse, error) {
	var resp tbrefund.TaobaorpreturngoodsagreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
