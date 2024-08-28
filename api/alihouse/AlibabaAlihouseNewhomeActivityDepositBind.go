package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivityDepositBind 销售活动绑定预存金商品id
// alibaba.alihouse.newhome.activity.deposit.bind
//
// 销售活动绑定预存金商品id
func AlibabaAlihouseNewhomeActivityDepositBind(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivityDepositBindAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeActivityDepositBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
