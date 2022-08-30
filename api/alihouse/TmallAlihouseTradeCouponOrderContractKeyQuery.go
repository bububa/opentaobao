package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderContractKeyQuery 查询电商券履约单合同key
// tmall.alihouse.trade.coupon.order.contract.key.query
//
// 查询电商券履约单合同地址
func TmallAlihouseTradeCouponOrderContractKeyQuery(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest, session string) (*alihouse.TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse, error) {
	var resp alihouse.TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
