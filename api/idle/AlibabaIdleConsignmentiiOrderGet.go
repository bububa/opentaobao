package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleConsignmentiiOrderGet 闲鱼寄卖V2订单查询
// alibaba.idle.consignmentii.order.get
//
// 闲鱼寄卖V2服务商以闲鱼交易买家身份查询订单信息
func AlibabaIdleConsignmentiiOrderGet(clt *core.SDKClient, req *idle.AlibabaIdleConsignmentiiOrderGetAPIRequest, session string) (*idle.AlibabaIdleConsignmentiiOrderGetAPIResponse, error) {
	var resp idle.AlibabaIdleConsignmentiiOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
