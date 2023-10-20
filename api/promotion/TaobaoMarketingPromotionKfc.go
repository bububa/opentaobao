package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaomarketingpromotionkfc 定向优惠活动名称与描述违禁词检查
// taobao.marketing.promotion.kfc
//
// 活动名称与描述违禁词检查
func Taobaomarketingpromotionkfc(clt *core.SDKClient, req *promotion.TaobaomarketingpromotionkfcAPIRequest, session string) (*promotion.TaobaomarketingpromotionkfcAPIResponse, error) {
	var resp promotion.TaobaomarketingpromotionkfcAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
