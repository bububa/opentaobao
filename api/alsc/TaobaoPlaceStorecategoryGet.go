package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoPlaceStorecategoryGet 获取门店类目信息
// taobao.place.storecategory.get
//
// 获取门店类目信息
func TaobaoPlaceStorecategoryGet(ctx context.Context, clt *core.SDKClient, req *alsc.TaobaoPlaceStorecategoryGetAPIRequest, resp *alsc.TaobaoPlaceStorecategoryGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
