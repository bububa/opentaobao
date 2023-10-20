package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscitemactivitydelete 删除无条件单品优惠活动
// taobao.promotionmisc.item.activity.delete
//
// 删除无条件单品优惠活动
func Taobaopromotionmiscitemactivitydelete(clt *core.SDKClient, req *promotion.TaobaopromotionmiscitemactivitydeleteAPIRequest, session string) (*promotion.TaobaopromotionmiscitemactivitydeleteAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscitemactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
