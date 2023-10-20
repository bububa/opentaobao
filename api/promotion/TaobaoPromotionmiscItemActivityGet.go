package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscitemactivityget 查询无条件单品优惠活动
// taobao.promotionmisc.item.activity.get
//
// 查询无条件单品优惠活动
func Taobaopromotionmiscitemactivityget(clt *core.SDKClient, req *promotion.TaobaopromotionmiscitemactivitygetAPIRequest, session string) (*promotion.TaobaopromotionmiscitemactivitygetAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscitemactivitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
