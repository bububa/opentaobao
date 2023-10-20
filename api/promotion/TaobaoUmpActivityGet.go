package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpactivityget 查询营销活动
// taobao.ump.activity.get
//
// 查询营销活动
func Taobaoumpactivityget(clt *core.SDKClient, req *promotion.TaobaoumpactivitygetAPIRequest, session string) (*promotion.TaobaoumpactivitygetAPIResponse, error) {
	var resp promotion.TaobaoumpactivitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
