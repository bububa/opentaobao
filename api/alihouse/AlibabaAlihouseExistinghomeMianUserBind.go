package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeMianUserBind 主账号入驻
// alibaba.alihouse.existinghome.mian.user.bind
//
// 主账号入驻
func AlibabaAlihouseExistinghomeMianUserBind(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeMianUserBindAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeMianUserBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
