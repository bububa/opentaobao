package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地址 APIRequest
taobao.degoperation.check.addr.status

激励
*/
type TaobaoDegoperationCheckAddrStatusRequest struct {
    model.Params

    // 奖品唯一标识
    sequenceNo   int64 

}

func NewTaobaoDegoperationCheckAddrStatusRequest() *TaobaoDegoperationCheckAddrStatusRequest{
    return &TaobaoDegoperationCheckAddrStatusRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDegoperationCheckAddrStatusRequest) GetApiMethodName() string {
    return "taobao.degoperation.check.addr.status"
}

func (r TaobaoDegoperationCheckAddrStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDegoperationCheckAddrStatusRequest) SetSequenceNo(sequenceNo int64) error {
    r.sequenceNo = sequenceNo
    r.Set("sequence_no", sequenceNo)
    return nil
}

func (r TaobaoDegoperationCheckAddrStatusRequest) GetSequenceNo() int64 {
    return r.sequenceNo
}

