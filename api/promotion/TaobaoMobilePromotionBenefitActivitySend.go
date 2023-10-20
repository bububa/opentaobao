package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaomobilepromotionbenefitactivitysend 手淘专用单用户发放接口
// taobao.mobile.promotion.benefit.activity.send
//
// 卖家活动中需要通过该API来发放对应的权益。手淘专用单用户发放接口。
func Taobaomobilepromotionbenefitactivitysend(clt *core.SDKClient, req *promotion.TaobaomobilepromotionbenefitactivitysendAPIRequest, session string) (*promotion.TaobaomobilepromotionbenefitactivitysendAPIResponse, error) {
	var resp promotion.TaobaomobilepromotionbenefitactivitysendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
