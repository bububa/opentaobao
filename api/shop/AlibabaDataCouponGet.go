package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Alibabadatacouponget 获取优惠券信息
// alibaba.data.coupon.get
//
// 获取优惠券信息，仅作客户端鉴权虚拟api使用
func Alibabadatacouponget(clt *core.SDKClient, req *shop.AlibabadatacoupongetAPIRequest, session string) (*shop.AlibabadatacoupongetAPIResponse, error) {
	var resp shop.AlibabadatacoupongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
