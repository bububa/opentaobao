package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscSaasCodecCodeAttrsQuery 码业务属性查询
// alibaba.alsc.saas.codec.code.attrs.query
//
// 码业务属性查询
func AlibabaAlscSaasCodecCodeAttrsQuery(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscSaasCodecCodeAttrsQueryAPIRequest, resp *alsc.AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
