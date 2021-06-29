package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消发货 API请求
taobao.rdc.aligenius.sendgoods.cancel

提供商家在仅退款中发送取消发货状态
*/
type TaobaoRdcAligeniusSendgoodsCancelRequest struct {
    model.Params
    // 请求参数
    _param   *CancelGoodsDTO
}

// 初始化TaobaoRdcAligeniusSendgoodsCancelRequest对象
func NewTaobaoRdcAligeniusSendgoodsCancelRequest() *TaobaoRdcAligeniusSendgoodsCancelRequest{
    return &TaobaoRdcAligeniusSendgoodsCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusSendgoodsCancelRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.sendgoods.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusSendgoodsCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *TaobaoRdcAligeniusSendgoodsCancelRequest) SetParam(_param *CancelGoodsDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoRdcAligeniusSendgoodsCancelRequest) GetParam() *CancelGoodsDTO {
    return r._param
}
