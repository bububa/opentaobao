package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmiscitemactivitylistget 查询无条件单品优惠活动列表
// taobao.promotionmisc.item.activity.list.get
//
// 查询无条件单品优惠活动列表
func Taobaopromotionmiscitemactivitylistget(clt *core.SDKClient, req *promotion.TaobaopromotionmiscitemactivitylistgetAPIRequest, session string) (*promotion.TaobaopromotionmiscitemactivitylistgetAPIResponse, error) {
	var resp promotion.TaobaopromotionmiscitemactivitylistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
