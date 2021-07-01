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
type YunosServiceCmnsCoaMessageCancelAPIRequest struct {
    model.Params
    // 消息ID
    _mid   int64
}

// 初始化YunosServiceCmnsCoaMessageCancelAPIRequest对象
func NewYunosServiceCmnsCoaMessageCancelRequest() *YunosServiceCmnsCoaMessageCancelAPIRequest{
    return &YunosServiceCmnsCoaMessageCancelAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageCancelAPIRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.cancel"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageCancelAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageCancelAPIRequest) SetMid(_mid int64) error {
    r._mid = _mid
    r.Set("mid", _mid)
    return nil
}

// Mid Getter
func (r YunosServiceCmnsCoaMessageCancelAPIRequest) GetMid() int64 {
    return r._mid
}
