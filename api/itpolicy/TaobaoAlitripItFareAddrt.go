package itpolicy

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/itpolicy"
)

/* 
【国际机票自有政策】单条往返添加 
taobao.alitrip.it.fare.addrt

自有政策往返添加接口
*/
func TaobaoAlitripItFareAddrt(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareAddrtAPIRequest, session string) (*itpolicy.TaobaoAlitripItFareAddrtAPIResponse, error) {
    var resp itpolicy.TaobaoAlitripItFareAddrtAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
