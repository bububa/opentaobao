package elife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/elife"
)

/* 
品牌惠卡券冲正退还 
taobao.elife.lifecard.refund

淘宝生活汇消费卡虚拟卡，线下冲正退货接口
*/
func TaobaoElifeLifecardRefund(clt *core.SDKClient, req *elife.TaobaoElifeLifecardRefundRequest, session string) (*elife.TaobaoElifeLifecardRefundAPIResponse, error) {
    var resp elife.TaobaoElifeLifecardRefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
