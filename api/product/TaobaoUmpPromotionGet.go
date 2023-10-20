package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaoumppromotionget 商品优惠详情查询
// taobao.ump.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
func Taobaoumppromotionget(clt *core.SDKClient, req *product.TaobaoumppromotiongetAPIRequest, session string) (*product.TaobaoumppromotiongetAPIResponse, error) {
	var resp product.TaobaoumppromotiongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
