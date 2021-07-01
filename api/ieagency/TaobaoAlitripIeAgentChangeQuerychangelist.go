package ieagency

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ieagency"
)

/* 
卖家查询改签列表 
taobao.alitrip.ie.agent.change.querychangelist

提供B2B卖家查看改签列表服务
*/
func TaobaoAlitripIeAgentChangeQuerychangelist(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest, session string) (*ieagency.TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse, error) {
    var resp ieagency.TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
