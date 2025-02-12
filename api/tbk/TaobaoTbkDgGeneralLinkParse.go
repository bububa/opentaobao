package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgGeneralLinkParse 淘宝客-推广者-万能解析
// taobao.tbk.dg.general.link.parse
//
// 淘宝客-推广者-万能转链
func TaobaoTbkDgGeneralLinkParse(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgGeneralLinkParseAPIRequest, resp *tbk.TaobaoTbkDgGeneralLinkParseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
