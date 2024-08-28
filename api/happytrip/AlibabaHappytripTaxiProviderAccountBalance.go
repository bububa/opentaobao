package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiProviderAccountBalance 供应商渠道余额
// alibaba.happytrip.taxi.provider.account.balance
//
// 查询不同供应商不同渠道账户余额
func AlibabaHappytripTaxiProviderAccountBalance(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiProviderAccountBalanceAPIRequest, resp *happytrip.AlibabaHappytripTaxiProviderAccountBalanceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
