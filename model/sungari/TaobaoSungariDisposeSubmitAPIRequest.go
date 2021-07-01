package sungari

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品商家处置提交任务 API请求
taobao.sungari.dispose.submit

商品商家处置信息接口，提供政府部门发送处置信息给阿里
*/
type TaobaoSungariDisposeSubmitAPIRequest struct {
    model.Params
    // 平台处置信息入参
    _info   *DisposeInfoDo
}

// 初始化TaobaoSungariDisposeSubmitAPIRequest对象
func NewTaobaoSungariDisposeSubmitRequest() *TaobaoSungariDisposeSubmitAPIRequest{
    return &TaobaoSungariDisposeSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSungariDisposeSubmitAPIRequest) GetApiMethodName() string {
    return "taobao.sungari.dispose.submit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSungariDisposeSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Info Setter
// 平台处置信息入参
func (r *TaobaoSungariDisposeSubmitAPIRequest) SetInfo(_info *DisposeInfoDo) error {
    r._info = _info
    r.Set("info", _info)
    return nil
}

// Info Getter
func (r TaobaoSungariDisposeSubmitAPIRequest) GetInfo() *DisposeInfoDo {
    return r._info
}
