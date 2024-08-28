package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplateDelete 删除运费模板
// taobao.delivery.template.delete
//
// 根据用户指定的模板ID删除指定的模板
func TaobaoDeliveryTemplateDelete(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplateDeleteAPIRequest, resp *tblogistics.TaobaoDeliveryTemplateDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
