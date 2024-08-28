package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplatesGet 获取用户下所有模板
// taobao.delivery.templates.get
//
// 根据用户ID获取用户下所有模板
func TaobaoDeliveryTemplatesGet(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplatesGetAPIRequest, resp *tblogistics.TaobaoDeliveryTemplatesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
