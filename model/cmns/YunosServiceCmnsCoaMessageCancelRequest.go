package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CMNS消息撤回 APIRequest
yunos.service.cmns.coa.message.cancel

此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。
*/
type YunosServiceCmnsCoaMessageCancelRequest struct {
    model.Params

    // 消息ID
    mid   int64 

}

func NewYunosServiceCmnsCoaMessageCancelRequest() *YunosServiceCmnsCoaMessageCancelRequest{
    return &YunosServiceCmnsCoaMessageCancelRequest{
        Params: model.NewParams(),
    }
}

func (r YunosServiceCmnsCoaMessageCancelRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.cancel"
}

func (r YunosServiceCmnsCoaMessageCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosServiceCmnsCoaMessageCancelRequest) SetMid(mid int64) error {
    r.mid = mid
    r.Set("mid", mid)
    return nil
}

func (r YunosServiceCmnsCoaMessageCancelRequest) GetMid() int64 {
    return r.mid
}

