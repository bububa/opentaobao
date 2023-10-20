package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionbenefitactivitysend 活动权益发放接口
// taobao.promotion.benefit.activity.send
//
// 活动权益发放接口，用于卖家针对活动进行权益发放
func Taobaopromotionbenefitactivitysend(clt *core.SDKClient, req *promotion.TaobaopromotionbenefitactivitysendAPIRequest, session string) (*promotion.TaobaopromotionbenefitactivitysendAPIResponse, error) {
	var resp promotion.TaobaopromotionbenefitactivitysendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
