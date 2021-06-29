package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
地址 API请求
taobao.degoperation.check.addr.status

激励
*/
type TaobaoDegoperationCheckAddrStatusRequest struct {
    model.Params
    // 奖品唯一标识
    sequenceNo   int64
}

// 初始化TaobaoDegoperationCheckAddrStatusRequest对象
func NewTaobaoDegoperationCheckAddrStatusRequest() *TaobaoDegoperationCheckAddrStatusRequest{
    return &TaobaoDegoperationCheckAddrStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationCheckAddrStatusRequest) GetApiMethodName() string {
    return "taobao.degoperation.check.addr.status"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationCheckAddrStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SequenceNo Setter
// 奖品唯一标识
func (r *TaobaoDegoperationCheckAddrStatusRequest) SetSequenceNo(sequenceNo int64) error {
    r.sequenceNo = sequenceNo
    r.Set("sequence_no", sequenceNo)
    return nil
}

// SequenceNo Getter
func (r TaobaoDegoperationCheckAddrStatusRequest) GetSequenceNo() int64 {
    return r.sequenceNo
}
