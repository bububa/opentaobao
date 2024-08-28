package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplateAdd 新增运费模板
// taobao.delivery.template.add
//
// 增加运费模板的外部接口
func TaobaoDeliveryTemplateAdd(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplateAddAPIRequest, resp *tblogistics.TaobaoDeliveryTemplateAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
