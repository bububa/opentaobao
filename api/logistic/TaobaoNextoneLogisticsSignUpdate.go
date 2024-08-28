package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoNextoneLogisticsSignUpdate AG物流签收状态写接口
// taobao.nextone.logistics.sign.update
//
// 商家上传退货的签收状态给AG
func TaobaoNextoneLogisticsSignUpdate(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoNextoneLogisticsSignUpdateAPIRequest, resp *logistic.TaobaoNextoneLogisticsSignUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
