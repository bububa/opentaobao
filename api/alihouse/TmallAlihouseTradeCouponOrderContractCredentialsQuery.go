package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderContractCredentialsQuery 查询用于电商券履约单合同下载的临时访问凭证
// tmall.alihouse.trade.coupon.order.contract.credentials.query
//
// 获取用于下载合同的临时aksk和安全token
func TmallAlihouseTradeCouponOrderContractCredentialsQuery(clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest, session string) (*alihouse.TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIResponse, error) {
	var resp alihouse.TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
