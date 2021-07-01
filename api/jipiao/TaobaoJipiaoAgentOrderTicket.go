package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

/* TaobaoJipiaoAgentOrderTicket
【机票代理商订单】订单回填票号/成功订单
taobao.jipiao.agent.order.ticket

淘宝机票代理商回填票号/成功订单 */
func TaobaoJipiaoAgentOrderTicket(clt *core.SDKClient, req *jipiao.TaobaoJipiaoAgentOrderTicketAPIRequest, session string) (*jipiao.TaobaoJipiaoAgentOrderTicketAPIResponse, error) {
	var resp jipiao.TaobaoJipiaoAgentOrderTicketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
