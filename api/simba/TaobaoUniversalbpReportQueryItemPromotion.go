package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryItemPromotion 宝贝主体报表查询
// taobao.universalbp.report.query.item.promotion
//
// 宝贝主体报表查询
func TaobaoUniversalbpReportQueryItemPromotion(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryItemPromotionAPIRequest, resp *simba.TaobaoUniversalbpReportQueryItemPromotionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
