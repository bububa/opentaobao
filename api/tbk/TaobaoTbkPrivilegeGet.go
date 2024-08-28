package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkPrivilegeGet 淘宝客-服务商-单品券高效转链
// taobao.tbk.privilege.get
//
// 单品券高效转链API
func TaobaoTbkPrivilegeGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkPrivilegeGetAPIRequest, resp *tbk.TaobaoTbkPrivilegeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
