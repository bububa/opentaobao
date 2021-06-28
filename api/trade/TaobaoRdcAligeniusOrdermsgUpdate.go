package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
订单消息状态回传 
taobao.rdc.aligenius.ordermsg.update

用于订单消息处理状态回传
*/
func TaobaoRdcAligeniusOrdermsgUpdate(clt *core.SDKClient, req *trade.TaobaoRdcAligeniusOrdermsgUpdateRequest, session string) (*trade.TaobaoRdcAligeniusOrdermsgUpdateAPIResponse, error) {
    var resp trade.TaobaoRdcAligeniusOrdermsgUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
