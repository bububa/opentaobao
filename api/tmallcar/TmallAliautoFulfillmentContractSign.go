package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoFulfillmentContractSign 合同签署
// tmall.aliauto.fulfillment.contract.sign
//
// 商家回传用户签署的合同信息
func TmallAliautoFulfillmentContractSign(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoFulfillmentContractSignAPIRequest, resp *tmallcar.TmallAliautoFulfillmentContractSignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
