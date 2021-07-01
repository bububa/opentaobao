package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
openmall服务地址区域码解析 
taobao.openmall.trade.address.parse

openmall服务，解析地址区域码，获取创建订单等接口中的区域码信息
*/
func TaobaoOpenmallTradeAddressParse(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeAddressParseAPIRequest, session string) (*openmall.TaobaoOpenmallTradeAddressParseAPIResponse, error) {
    var resp openmall.TaobaoOpenmallTradeAddressParseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
