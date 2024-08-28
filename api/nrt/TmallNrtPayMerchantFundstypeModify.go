package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtPayMerchantFundstypeModify 修改摊位分账类型
// tmall.nrt.pay.merchant.fundstype.modify
//
// 修改摊位分账类型
func TmallNrtPayMerchantFundstypeModify(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtPayMerchantFundstypeModifyAPIRequest, resp *nrt.TmallNrtPayMerchantFundstypeModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
