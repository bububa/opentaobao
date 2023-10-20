package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobustvmcreateorderset 线下自助机创建订单
// taobao.bus.tvmcreateorder.set
//
// 提供给汽车票线下自助机的创建订单使用
func Taobaobustvmcreateorderset(clt *core.SDKClient, req *bus.TaobaobustvmcreateordersetAPIRequest, session string) (*bus.TaobaobustvmcreateordersetAPIResponse, error) {
	var resp bus.TaobaobustvmcreateordersetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
