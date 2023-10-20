package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

// Taobaolarkiotordergetgoodslist iot渠道获取卖品信息
// taobao.lark.iot.order.getgoodslist
//
// iot无人超市服务商通过接口获取影院的可售卖品数据
func Taobaolarkiotordergetgoodslist(clt *core.SDKClient, req *larkiot.TaobaolarkiotordergetgoodslistAPIRequest, session string) (*larkiot.TaobaolarkiotordergetgoodslistAPIResponse, error) {
	var resp larkiot.TaobaolarkiotordergetgoodslistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
