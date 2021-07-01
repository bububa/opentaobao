package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
廉航辅营产品投放 
taobao.alitrip.totoro.auxproduct.push

廉航辅营产品投放接口
*/
func TaobaoAlitripTotoroAuxproductPush(clt *core.SDKClient, req *flight.TaobaoAlitripTotoroAuxproductPushAPIRequest, session string) (*flight.TaobaoAlitripTotoroAuxproductPushAPIResponse, error) {
    var resp flight.TaobaoAlitripTotoroAuxproductPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
