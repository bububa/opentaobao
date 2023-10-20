package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpactivityadd 新增优惠活动
// taobao.ump.activity.add
//
// 新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
func Taobaoumpactivityadd(clt *core.SDKClient, req *promotion.TaobaoumpactivityaddAPIRequest, session string) (*promotion.TaobaoumpactivityaddAPIResponse, error) {
	var resp promotion.TaobaoumpactivityaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
