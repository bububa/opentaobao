package refund

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/refund"
)

/* 
查询退款留言/凭证列表 
taobao.refund.messages.get

查询退款留言/凭证列表
*/
func TaobaoRefundMessagesGet(clt *core.SDKClient, req *refund.TaobaoRefundMessagesGetRequest, session string) (*refund.TaobaoRefundMessagesGetResponse, error) {
    var resp refund.TaobaoRefundMessagesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
