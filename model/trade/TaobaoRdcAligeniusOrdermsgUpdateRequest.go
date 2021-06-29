package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单消息状态回传 API请求
taobao.rdc.aligenius.ordermsg.update

用于订单消息处理状态回传
*/
type TaobaoRdcAligeniusOrdermsgUpdateRequest struct {
    model.Params
    // 子订单（消息中传的子订单）
    oid   int64
    // 处理状态，1=成功，2=处理失败
    status   int64
    // 主订单（消息中传的主订单）
    tid   int64
}

// 初始化TaobaoRdcAligeniusOrdermsgUpdateRequest对象
func NewTaobaoRdcAligeniusOrdermsgUpdateRequest() *TaobaoRdcAligeniusOrdermsgUpdateRequest{
    return &TaobaoRdcAligeniusOrdermsgUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.ordermsg.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Oid Setter
// 子订单（消息中传的子订单）
func (r *TaobaoRdcAligeniusOrdermsgUpdateRequest) SetOid(oid int64) error {
    r.oid = oid
    r.Set("oid", oid)
    return nil
}

// Oid Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetOid() int64 {
    return r.oid
}
// Status Setter
// 处理状态，1=成功，2=处理失败
func (r *TaobaoRdcAligeniusOrdermsgUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetStatus() int64 {
    return r.status
}
// Tid Setter
// 主订单（消息中传的主订单）
func (r *TaobaoRdcAligeniusOrdermsgUpdateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetTid() int64 {
    return r.tid
}
