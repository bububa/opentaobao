package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Tmalldeviceitempromotionget 智能硬件上商品优惠获取
// tmall.device.item.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
func Tmalldeviceitempromotionget(clt *core.SDKClient, req *iot.TmalldeviceitempromotiongetAPIRequest, session string) (*iot.TmalldeviceitempromotiongetAPIResponse, error) {
	var resp iot.TmalldeviceitempromotiongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
