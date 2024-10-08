package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomOrderPreauthorizeCreate 业务办理结果
// alibaba.alicom.order.preauthorize.create
//
// 授授权:签约结果通知
func AlibabaAlicomOrderPreauthorizeCreate(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAlicomOrderPreauthorizeCreateAPIRequest, resp *alicom.AlibabaAlicomOrderPreauthorizeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
