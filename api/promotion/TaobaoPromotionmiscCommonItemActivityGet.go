package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmisccommonitemactivityget 查询通用单品优惠活动
// taobao.promotionmisc.common.item.activity.get
//
// 查询通用单品优惠活动。
func Taobaopromotionmisccommonitemactivityget(clt *core.SDKClient, req *promotion.TaobaopromotionmisccommonitemactivitygetAPIRequest, session string) (*promotion.TaobaopromotionmisccommonitemactivitygetAPIResponse, error) {
	var resp promotion.TaobaopromotionmisccommonitemactivitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
