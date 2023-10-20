package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketetickettimeexpand 订单延时接口
// taobao.vmarket.eticket.time.expand
//
// 提供码商操作订单延期接口
func Taobaovmarketetickettimeexpand(clt *core.SDKClient, req *eticket.TaobaovmarketetickettimeexpandAPIRequest, session string) (*eticket.TaobaovmarketetickettimeexpandAPIResponse, error) {
	var resp eticket.TaobaovmarketetickettimeexpandAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
