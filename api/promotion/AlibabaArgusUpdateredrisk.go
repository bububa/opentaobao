package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaArgusUpdateredrisk 更新红线价格
// alibaba.argus.updateredrisk
//
// 商品健康中心新增红线价格规则
func AlibabaArgusUpdateredrisk(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaArgusUpdateredriskAPIRequest, resp *promotion.AlibabaArgusUpdateredriskAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
