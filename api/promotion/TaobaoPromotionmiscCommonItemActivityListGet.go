package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmisccommonitemactivitylistget 查询通用单品优惠活动列表
// taobao.promotionmisc.common.item.activity.list.get
//
// 查询通用单品优惠活动列表。
func Taobaopromotionmisccommonitemactivitylistget(clt *core.SDKClient, req *promotion.TaobaopromotionmisccommonitemactivitylistgetAPIRequest, session string) (*promotion.TaobaopromotionmisccommonitemactivitylistgetAPIResponse, error) {
	var resp promotion.TaobaopromotionmisccommonitemactivitylistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
