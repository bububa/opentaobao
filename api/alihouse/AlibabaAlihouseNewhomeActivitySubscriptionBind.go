package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivitySubscriptionBind 销售活动绑定认购商品
// alibaba.alihouse.newhome.activity.subscription.bind
//
// 销售活动绑定认购商品
func AlibabaAlihouseNewhomeActivitySubscriptionBind(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
