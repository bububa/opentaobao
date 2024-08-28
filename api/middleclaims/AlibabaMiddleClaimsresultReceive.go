package middleclaims

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/middleclaims"
)

// AlibabaMiddleClaimsresultReceive 国际化中台服务域接收理赔结果
// alibaba.middle.claimsresult.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔结果的处理后，将该结果返回至服务域
func AlibabaMiddleClaimsresultReceive(ctx context.Context, clt *core.SDKClient, req *middleclaims.AlibabaMiddleClaimsresultReceiveAPIRequest, resp *middleclaims.AlibabaMiddleClaimsresultReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
