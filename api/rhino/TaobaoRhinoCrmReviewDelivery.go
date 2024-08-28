package rhino

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// TaobaoRhinoCrmReviewDelivery crm实体预询期
// taobao.rhino.crm.review.delivery
//
// crm实体预询期
func TaobaoRhinoCrmReviewDelivery(ctx context.Context, clt *core.SDKClient, req *rhino.TaobaoRhinoCrmReviewDeliveryAPIRequest, resp *rhino.TaobaoRhinoCrmReviewDeliveryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
