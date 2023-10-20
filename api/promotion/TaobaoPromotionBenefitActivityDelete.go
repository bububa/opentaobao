package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionbenefitactivitydelete 删除关联的活动权益
// taobao.promotion.benefit.activity.delete
//
// 删除关联的活动权益
func Taobaopromotionbenefitactivitydelete(clt *core.SDKClient, req *promotion.TaobaopromotionbenefitactivitydeleteAPIRequest, session string) (*promotion.TaobaopromotionbenefitactivitydeleteAPIResponse, error) {
	var resp promotion.TaobaopromotionbenefitactivitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
