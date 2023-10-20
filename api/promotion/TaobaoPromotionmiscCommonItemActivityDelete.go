package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionmisccommonitemactivitydelete 删除通用单品优惠活动
// taobao.promotionmisc.common.item.activity.delete
//
// 删除通用单品优惠活动。
func Taobaopromotionmisccommonitemactivitydelete(clt *core.SDKClient, req *promotion.TaobaopromotionmisccommonitemactivitydeleteAPIRequest, session string) (*promotion.TaobaopromotionmisccommonitemactivitydeleteAPIResponse, error) {
	var resp promotion.TaobaopromotionmisccommonitemactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
