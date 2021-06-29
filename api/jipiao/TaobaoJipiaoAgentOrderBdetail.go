package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
【机票代理商订单】采购订单详情 
taobao.jipiao.agent.order.bdetail

根据淘宝系统订单号获取订单详情信息
*/
func TaobaoJipiaoAgentOrderBdetail(clt *core.SDKClient, req *jipiao.TaobaoJipiaoAgentOrderBdetailRequest, session string) (*jipiao.TaobaoJipiaoAgentOrderBdetailAPIResponse, error) {
    var resp jipiao.TaobaoJipiaoAgentOrderBdetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
