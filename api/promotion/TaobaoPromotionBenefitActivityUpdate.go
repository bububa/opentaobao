package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionbenefitactivityupdate 修改关联的活动权益
// taobao.promotion.benefit.activity.update
//
// 修改卖家活动中关联的对应的权益。
func Taobaopromotionbenefitactivityupdate(clt *core.SDKClient, req *promotion.TaobaopromotionbenefitactivityupdateAPIRequest, session string) (*promotion.TaobaopromotionbenefitactivityupdateAPIResponse, error) {
	var resp promotion.TaobaopromotionbenefitactivityupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
