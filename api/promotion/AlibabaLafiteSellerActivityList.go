package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabalafiteselleractivitylist 商家自运营活动列表
// alibaba.lafite.seller.activity.list
//
// 商家查询自己配置的活动列表
func Alibabalafiteselleractivitylist(clt *core.SDKClient, req *promotion.AlibabalafiteselleractivitylistAPIRequest, session string) (*promotion.AlibabalafiteselleractivitylistAPIResponse, error) {
	var resp promotion.AlibabalafiteselleractivitylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
