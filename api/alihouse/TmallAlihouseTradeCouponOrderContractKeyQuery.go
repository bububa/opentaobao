package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Tmallalihousetradecouponordercontractkeyquery 查询电商券履约单合同key
// tmall.alihouse.trade.coupon.order.contract.key.query
//
// 查询电商券履约单合同地址
func Tmallalihousetradecouponordercontractkeyquery(clt *core.SDKClient, req *alihouse.TmallalihousetradecouponordercontractkeyqueryAPIRequest, session string) (*alihouse.TmallalihousetradecouponordercontractkeyqueryAPIResponse, error) {
	var resp alihouse.TmallalihousetradecouponordercontractkeyqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
