package consignplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/consignplatform"
)

// Cainiaoconsignplatformordercreate 菜鸟发货工作台创建订单
// cainiao.consignplatform.order.create
//
// 菜鸟发货工作台，商家或者isv通过api进行订单写入操作
func Cainiaoconsignplatformordercreate(clt *core.SDKClient, req *consignplatform.CainiaoconsignplatformordercreateAPIRequest, session string) (*consignplatform.CainiaoconsignplatformordercreateAPIResponse, error) {
	var resp consignplatform.CainiaoconsignplatformordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
