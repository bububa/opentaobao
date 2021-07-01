package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
批量获取openmall退款单 
taobao.openmall.refund.batch.get

批量获取openmall退款单
注意：该接口信息存在延迟，如需实时详情请访问taobao.openmall.refund.get
*/
func TaobaoOpenmallRefundBatchGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundBatchGetAPIRequest, session string) (*openmall.TaobaoOpenmallRefundBatchGetAPIResponse, error) {
    var resp openmall.TaobaoOpenmallRefundBatchGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
