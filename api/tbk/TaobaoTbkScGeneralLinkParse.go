package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScGeneralLinkParse 淘宝客-服务商-万能解析
// taobao.tbk.sc.general.link.parse
//
// 淘宝客-推广者-万能转链
func TaobaoTbkScGeneralLinkParse(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScGeneralLinkParseAPIRequest, resp *tbk.TaobaoTbkScGeneralLinkParseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
