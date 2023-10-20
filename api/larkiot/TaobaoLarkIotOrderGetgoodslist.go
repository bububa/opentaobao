package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

// TaobaoLarkIotOrderGetgoodslist iot渠道获取卖品信息
// taobao.lark.iot.order.getgoodslist
//
// iot无人超市服务商通过接口获取影院的可售卖品数据
func TaobaoLarkIotOrderGetgoodslist(clt *core.SDKClient, req *larkiot.TaobaoLarkIotOrderGetgoodslistAPIRequest, session string) (*larkiot.TaobaoLarkIotOrderGetgoodslistAPIResponse, error) {
	var resp larkiot.TaobaoLarkIotOrderGetgoodslistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
