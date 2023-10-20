package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionbenefitactivityrelation 关联活动权益
// taobao.promotion.benefit.activity.relation
//
// 卖家活动中需要通过该API来关联的对应的权益。
func Taobaopromotionbenefitactivityrelation(clt *core.SDKClient, req *promotion.TaobaopromotionbenefitactivityrelationAPIRequest, session string) (*promotion.TaobaopromotionbenefitactivityrelationAPIResponse, error) {
	var resp promotion.TaobaopromotionbenefitactivityrelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
