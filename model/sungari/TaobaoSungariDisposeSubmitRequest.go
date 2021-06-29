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
type TaobaoSungariDisposeSubmitRequest struct {
    model.Params
    // 平台处置信息入参
    info   *DisposeInfoDo
}

// 初始化TaobaoSungariDisposeSubmitRequest对象
func NewTaobaoSungariDisposeSubmitRequest() *TaobaoSungariDisposeSubmitRequest{
    return &TaobaoSungariDisposeSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSungariDisposeSubmitRequest) GetApiMethodName() string {
    return "taobao.sungari.dispose.submit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSungariDisposeSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Info Setter
// 平台处置信息入参
func (r *TaobaoSungariDisposeSubmitRequest) SetInfo(info *DisposeInfoDo) error {
    r.info = info
    r.Set("info", info)
    return nil
}

// Info Getter
func (r TaobaoSungariDisposeSubmitRequest) GetInfo() *DisposeInfoDo {
    return r.info
}
