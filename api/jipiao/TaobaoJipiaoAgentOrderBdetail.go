package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaojipiaoagentorderbdetail 【机票代理商订单】采购订单详情
// taobao.jipiao.agent.order.bdetail
//
// 根据淘宝系统订单号获取订单详情信息
func Taobaojipiaoagentorderbdetail(clt *core.SDKClient, req *jipiao.TaobaojipiaoagentorderbdetailAPIRequest, session string) (*jipiao.TaobaojipiaoagentorderbdetailAPIResponse, error) {
	var resp jipiao.TaobaojipiaoagentorderbdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
