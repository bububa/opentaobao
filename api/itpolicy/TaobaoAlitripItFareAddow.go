package itpolicy

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/itpolicy"
)

/* 
【国际机票自有政策】单条单程添加 
taobao.alitrip.it.fare.addow

自有政策单程添加接口，重复的老数据会被删除，重复判断规则同excel
*/
func TaobaoAlitripItFareAddow(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareAddowAPIRequest, session string) (*itpolicy.TaobaoAlitripItFareAddowAPIResponse, error) {
    var resp itpolicy.TaobaoAlitripItFareAddowAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
