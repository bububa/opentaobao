package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoWlbStoresBaseinfoGet 商家按照仓的类型获取仓库接口
// taobao.wlb.stores.baseinfo.get
//
// 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
func TaobaoWlbStoresBaseinfoGet(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoWlbStoresBaseinfoGetAPIRequest, resp *tblogistics.TaobaoWlbStoresBaseinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
