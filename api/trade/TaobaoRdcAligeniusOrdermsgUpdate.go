package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoRdcAligeniusOrdermsgUpdate 订单消息状态回传
// taobao.rdc.aligenius.ordermsg.update
//
// 用于订单消息处理状态回传
func TaobaoRdcAligeniusOrdermsgUpdate(clt *core.SDKClient, req *trade.TaobaoRdcAligeniusOrdermsgUpdateAPIRequest, resp *trade.TaobaoRdcAligeniusOrdermsgUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
