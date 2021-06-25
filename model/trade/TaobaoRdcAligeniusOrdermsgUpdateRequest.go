package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单消息状态回传 APIRequest
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

func NewTaobaoRdcAligeniusOrdermsgUpdateRequest() *TaobaoRdcAligeniusOrdermsgUpdateRequest{
    return &TaobaoRdcAligeniusOrdermsgUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.ordermsg.update"
}

func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdcAligeniusOrdermsgUpdateRequest) SetOid(oid int64) error {
    r.oid = oid
    r.Set("oid", oid)
    return nil
}

func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetOid() int64 {
    return r.oid
}

func (r *TaobaoRdcAligeniusOrdermsgUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoRdcAligeniusOrdermsgUpdateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoRdcAligeniusOrdermsgUpdateRequest) GetTid() int64 {
    return r.tid
}

