package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryItemPromotion 宝贝主体报表查询
// taobao.universalbp.report.query.item.promotion
//
// 宝贝主体报表查询
func TaobaoUniversalbpReportQueryItemPromotion(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryItemPromotionAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryItemPromotionAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryItemPromotionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
