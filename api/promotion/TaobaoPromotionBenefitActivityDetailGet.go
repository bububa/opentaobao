package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaopromotionbenefitactivitydetailget 活动关联的权益详情获取
// taobao.promotion.benefit.activity.detail.get
//
// 活动关联的权益详情获取
func Taobaopromotionbenefitactivitydetailget(clt *core.SDKClient, req *promotion.TaobaopromotionbenefitactivitydetailgetAPIRequest, session string) (*promotion.TaobaopromotionbenefitactivitydetailgetAPIResponse, error) {
	var resp promotion.TaobaopromotionbenefitactivitydetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
