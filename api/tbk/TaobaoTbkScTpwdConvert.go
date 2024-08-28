package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScTpwdConvert 淘宝客-服务商-淘口令解析&转链
// taobao.tbk.sc.tpwd.convert
//
// 支持通过淘口令解析商品id，并提供对应的淘客转链接
func TaobaoTbkScTpwdConvert(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScTpwdConvertAPIRequest, resp *tbk.TaobaoTbkScTpwdConvertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
