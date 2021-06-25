package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消发货 APIRequest
taobao.rdc.aligenius.sendgoods.cancel

提供商家在仅退款中发送取消发货状态
*/
type TaobaoRdcAligeniusSendgoodsCancelRequest struct {
    model.Params

    // 请求参数
    param   *CancelGoodsDto 

}

func NewTaobaoRdcAligeniusSendgoodsCancelRequest() *TaobaoRdcAligeniusSendgoodsCancelRequest{
    return &TaobaoRdcAligeniusSendgoodsCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusSendgoodsCancelRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.sendgoods.cancel"
}

func (r TaobaoRdcAligeniusSendgoodsCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdcAligeniusSendgoodsCancelRequest) SetParam(param *CancelGoodsDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoRdcAligeniusSendgoodsCancelRequest) GetParam() *CancelGoodsDto {
    return r.param
}

