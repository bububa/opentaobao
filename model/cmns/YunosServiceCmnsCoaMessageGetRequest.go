package cmns

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息详情查询 APIRequest
yunos.service.cmns.coa.message.get

第三方应用开发者调用此接口查询消息详情，只能查询此appKey发的消息
*/
type YunosServiceCmnsCoaMessageGetRequest struct {
    model.Params

    // 消息id
    mid   int64 

}

func NewYunosServiceCmnsCoaMessageGetRequest() *YunosServiceCmnsCoaMessageGetRequest{
    return &YunosServiceCmnsCoaMessageGetRequest{
        Params: model.NewParams(),
    }
}

func (r YunosServiceCmnsCoaMessageGetRequest) GetApiMethodName() string {
    return "yunos.service.cmns.coa.message.get"
}

func (r YunosServiceCmnsCoaMessageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosServiceCmnsCoaMessageGetRequest) SetMid(mid int64) error {
    r.mid = mid
    r.Set("mid", mid)
    return nil
}

func (r YunosServiceCmnsCoaMessageGetRequest) GetMid() int64 {
    return r.mid
}

