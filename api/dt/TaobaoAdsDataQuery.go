package dt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// TaobaoAdsDataQuery 导入数据查询
// taobao.ads.data.query
//
// 导入数据查询
func TaobaoAdsDataQuery(ctx context.Context, clt *core.SDKClient, req *dt.TaobaoAdsDataQueryAPIRequest, resp *dt.TaobaoAdsDataQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
