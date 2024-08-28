package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeitemFeaturesSync 同步品活动标
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
func AlibabaAlihouseNewhomeTradeitemFeaturesSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
