package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgGeneralLinkConvert 淘宝客-推广者-万能转链
// taobao.tbk.dg.general.link.convert
//
// 淘宝客-推广者-万能转链
func TaobaoTbkDgGeneralLinkConvert(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgGeneralLinkConvertAPIRequest, resp *tbk.TaobaoTbkDgGeneralLinkConvertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
