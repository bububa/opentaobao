package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Tmallalihousetradecouponordercodeexchange 核销券码
// tmall.alihouse.trade.coupon.order.code.exchange
//
// ETC核销券码
func Tmallalihousetradecouponordercodeexchange(clt *core.SDKClient, req *alihouse.TmallalihousetradecouponordercodeexchangeAPIRequest, session string) (*alihouse.TmallalihousetradecouponordercodeexchangeAPIResponse, error) {
	var resp alihouse.TmallalihousetradecouponordercodeexchangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
