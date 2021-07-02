package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// TaobaoJipiaoAgentOrderSearch 【机票代理商订单】订单搜索
// taobao.jipiao.agent.order.search
//
// 卖家根据条件查询淘宝订单id列表
func TaobaoJipiaoAgentOrderSearch(clt *core.SDKClient, req *jipiao.TaobaoJipiaoAgentOrderSearchAPIRequest, session string) (*jipiao.TaobaoJipiaoAgentOrderSearchAPIResponse, error) {
	var resp jipiao.TaobaoJipiaoAgentOrderSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
