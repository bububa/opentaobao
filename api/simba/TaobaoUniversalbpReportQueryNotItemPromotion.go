package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryNotItemPromotion 其他主体报表查询
// taobao.universalbp.report.query.not.item.promotion
//
// 其他主体报表查询
func TaobaoUniversalbpReportQueryNotItemPromotion(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryNotItemPromotionAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
