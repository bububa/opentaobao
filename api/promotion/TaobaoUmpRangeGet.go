package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumprangeget 查询活动范围
// taobao.ump.range.get
//
// 查询某个卖家所有参加或者不参加某项活动的物品
func Taobaoumprangeget(clt *core.SDKClient, req *promotion.TaobaoumprangegetAPIRequest, session string) (*promotion.TaobaoumprangegetAPIResponse, error) {
	var resp promotion.TaobaoumprangegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
