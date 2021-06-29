package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CMNS消息撤回 API请求
yunos.service.cmns.coa.message.cancel

此接口用户撤回之前已经发出去的消息，根据消息ID撤回，只能撤回此appKey创建的消息。
*/
type YunosServiceCmnsCoaMessageCancelRequest struct {
    model.Params
    // 消息ID
    mid   int64
}

// 初始化YunosServiceCmnsCoaMessageCancelRequest对象
func NewYunosServiceCmnsCoaMessageCancelRequest() *YunosServiceCmnsCoaMessageCancelRequest{
    return &YunosServiceCmnsCoaMessageCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageCancelRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.cancel"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageCancelRequest) SetMid(mid int64) error {
    r.mid = mid
    r.Set("mid", mid)
    return nil
}

// Mid Getter
func (r YunosServiceCmnsCoaMessageCancelRequest) GetMid() int64 {
    return r.mid
}
