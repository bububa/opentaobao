package caipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// TaobaoCaipiaoPresentStatGet 获取卖家按天统计的彩票赠送数据
// taobao.caipiao.present.stat.get
//
// 查询卖家一段时间内按天统计的彩票赠送数据，只支持查询90天以内的数据.
func TaobaoCaipiaoPresentStatGet(ctx context.Context, clt *core.SDKClient, req *caipiao.TaobaoCaipiaoPresentStatGetAPIRequest, resp *caipiao.TaobaoCaipiaoPresentStatGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
