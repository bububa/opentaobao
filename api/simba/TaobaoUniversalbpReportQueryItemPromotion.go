package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryItemPromotion 宝贝主体报表查询
// taobao.universalbp.report.query.item.promotion
//
// 宝贝主体报表查询
func TaobaoUniversalbpReportQueryItemPromotion(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryItemPromotionAPIRequest, resp *simba.TaobaoUniversalbpReportQueryItemPromotionAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
