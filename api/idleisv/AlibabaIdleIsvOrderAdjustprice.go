package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleisvorderadjustprice 闲鱼服务商订单价格修改接口
// alibaba.idle.isv.order.adjustprice
//
// 闲鱼用户通过授权的服务商修改订单价格和邮费
func Alibabaidleisvorderadjustprice(clt *core.SDKClient, req *idleisv.AlibabaidleisvorderadjustpriceAPIRequest, session string) (*idleisv.AlibabaidleisvorderadjustpriceAPIResponse, error) {
	var resp idleisv.AlibabaidleisvorderadjustpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
