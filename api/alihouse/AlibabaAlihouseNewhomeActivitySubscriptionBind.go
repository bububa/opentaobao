package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivitySubscriptionBind 销售活动绑定认购商品
// alibaba.alihouse.newhome.activity.subscription.bind
//
// 销售活动绑定认购商品
func AlibabaAlihouseNewhomeActivitySubscriptionBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivitySubscriptionBindAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeActivitySubscriptionBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
