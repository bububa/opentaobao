package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CMNS消息发送到达查询 API请求
yunos.service.cmns.coa.messageresult.get

CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
*/
type YunosServiceCmnsCoaMessageresultGetRequest struct {
    model.Params
    // 消息ID
    mid   int64
}

// 初始化YunosServiceCmnsCoaMessageresultGetRequest对象
func NewYunosServiceCmnsCoaMessageresultGetRequest() *YunosServiceCmnsCoaMessageresultGetRequest{
    return &YunosServiceCmnsCoaMessageresultGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaMessageresultGetRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.messageresult.get"
}

// IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaMessageresultGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mid Setter
// 消息ID
func (r *YunosServiceCmnsCoaMessageresultGetRequest) SetMid(mid int64) error {
    r.mid = mid
    r.Set("mid", mid)
    return nil
}

// Mid Getter
func (r YunosServiceCmnsCoaMessageresultGetRequest) GetMid() int64 {
    return r.mid
}
