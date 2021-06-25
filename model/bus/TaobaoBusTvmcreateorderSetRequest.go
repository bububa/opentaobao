package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机创建订单 APIRequest
taobao.bus.tvmcreateorder.set

提供给汽车票线下自助机的创建订单使用
*/
type TaobaoBusTvmcreateorderSetRequest struct {
    model.Params

    // 创建订单对象
    paramTVMCreateOrderRQ   *TvmCreateOrderRq 

}

func NewTaobaoBusTvmcreateorderSetRequest() *TaobaoBusTvmcreateorderSetRequest{
    return &TaobaoBusTvmcreateorderSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusTvmcreateorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmcreateorder.set"
}

func (r TaobaoBusTvmcreateorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusTvmcreateorderSetRequest) SetParamTVMCreateOrderRQ(paramTVMCreateOrderRQ *TvmCreateOrderRq) error {
    r.paramTVMCreateOrderRQ = paramTVMCreateOrderRQ
    r.Set("param_t_v_m_create_order_r_q", paramTVMCreateOrderRQ)
    return nil
}

func (r TaobaoBusTvmcreateorderSetRequest) GetParamTVMCreateOrderRQ() *TvmCreateOrderRq {
    return r.paramTVMCreateOrderRQ
}

