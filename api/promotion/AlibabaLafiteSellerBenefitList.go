package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabalafitesellerbenefitlist 商家自运营权益列表
// alibaba.lafite.seller.benefit.list
//
// 小程序isv可使用该接口获取权益列表
func Alibabalafitesellerbenefitlist(clt *core.SDKClient, req *promotion.AlibabalafitesellerbenefitlistAPIRequest, session string) (*promotion.AlibabalafitesellerbenefitlistAPIResponse, error) {
	var resp promotion.AlibabalafitesellerbenefitlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
