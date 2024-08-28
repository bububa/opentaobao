package usergrowth

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoUsergrowthDhhDeliveryAsk 广告曝光前判定接口V2
// taobao.usergrowth.dhh.delivery.ask
//
// 提供给媒体在曝光广告前调用
func TaobaoUsergrowthDhhDeliveryAsk(ctx context.Context, clt *core.SDKClient, req *usergrowth.TaobaoUsergrowthDhhDeliveryAskAPIRequest, resp *usergrowth.TaobaoUsergrowthDhhDeliveryAskAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
