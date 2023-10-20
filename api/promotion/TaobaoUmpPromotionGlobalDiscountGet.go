package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumppromotionglobaldiscountget 获取卖家最低折扣
// taobao.ump.promotion.global.discount.get
//
// 提供卖家最低折扣查询功能
func Taobaoumppromotionglobaldiscountget(clt *core.SDKClient, req *promotion.TaobaoumppromotionglobaldiscountgetAPIRequest, session string) (*promotion.TaobaoumppromotionglobaldiscountgetAPIResponse, error) {
	var resp promotion.TaobaoumppromotionglobaldiscountgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
