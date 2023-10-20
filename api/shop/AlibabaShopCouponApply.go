package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Alibabashopcouponapply 通用店铺券领券接口
// alibaba.shop.coupon.apply
//
// 店铺小部件和模块开发的isv通用店铺券领券接口
func Alibabashopcouponapply(clt *core.SDKClient, req *shop.AlibabashopcouponapplyAPIRequest, session string) (*shop.AlibabashopcouponapplyAPIResponse, error) {
	var resp shop.AlibabashopcouponapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
