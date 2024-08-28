package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeActivityDepositUnbind 销售活动解绑预存金商品
// alibaba.alihouse.newhome.activity.deposit.unbind
//
// 销售活动解绑预存金商品
func AlibabaAlihouseNewhomeActivityDepositUnbind(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeActivityDepositUnbindAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeActivityDepositUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
