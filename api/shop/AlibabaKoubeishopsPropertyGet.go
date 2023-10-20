package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Alibabakoubeishopspropertyget 口碑店铺列表推荐
// alibaba.koubeishops.property.get
//
// 推荐用户附近的美食门店
func Alibabakoubeishopspropertyget(clt *core.SDKClient, req *shop.AlibabakoubeishopspropertygetAPIRequest, session string) (*shop.AlibabakoubeishopspropertygetAPIResponse, error) {
	var resp shop.AlibabakoubeishopspropertygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
