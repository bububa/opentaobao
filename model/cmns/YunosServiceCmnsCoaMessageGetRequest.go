package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息详情查询 API请求
yunos.service.cmns.coa.message.get

第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息
*/
type YunosServiceCmnsCoaMessageGetRequest struct {
    model.Params
    // 消息id
    _mid   int64
}

// 初始化YunosServiceCmnsCoaMessageGetRequest对象
func NewYunosServiceCmnsCoaMessageGetRequest() *YunosServiceCmnsCoaMessageGetRequest{
    return &YunosServiceCmnsCoaMessageGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageGetRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.get"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mid Setter
// 消息id
func (r *YunosServiceCmnsCoaMessageGetRequest) SetMid(_mid int64) error {
    r._mid = _mid
    r.Set("mid", _mid)
    return nil
}

// Mid Getter
func (r YunosServiceCmnsCoaMessageGetRequest) GetMid() int64 {
    return r._mid
}
