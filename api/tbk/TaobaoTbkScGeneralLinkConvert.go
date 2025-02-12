package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScGeneralLinkConvert 淘宝客-服务商-万能转链
// taobao.tbk.sc.general.link.convert
//
// 淘宝客-服务商-万能转链
func TaobaoTbkScGeneralLinkConvert(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScGeneralLinkConvertAPIRequest, resp *tbk.TaobaoTbkScGeneralLinkConvertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
