package qt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// TaobaoTsSubscribeGet 淘宝服务订购关系查询
// taobao.ts.subscribe.get
//
// ts订购关系状态查询. 暂只支持1口价服务.
func TaobaoTsSubscribeGet(ctx context.Context, clt *core.SDKClient, req *qt.TaobaoTsSubscribeGetAPIRequest, resp *qt.TaobaoTsSubscribeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
