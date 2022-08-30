package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeOrderdetailEntryGet 获取订单详情页跳转入口配置
// taobao.opentrade.orderdetail.entry.get
//
// 获取订单详情页跳转入口配置
func TaobaoOpentradeOrderdetailEntryGet(clt *core.SDKClient, req *opentrade.TaobaoOpentradeOrderdetailEntryGetAPIRequest, session string) (*opentrade.TaobaoOpentradeOrderdetailEntryGetAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeOrderdetailEntryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
