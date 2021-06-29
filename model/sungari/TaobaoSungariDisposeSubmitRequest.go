package sungari

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品商家处置提交任务 APIRequest
taobao.sungari.dispose.submit

商品商家处置信息接口，提供政府部门发送处置信息给阿里
*/
type TaobaoSungariDisposeSubmitRequest struct {
    model.Params

    // 平台处置信息入参
    info   *DisposeInfoDo 

}

func NewTaobaoSungariDisposeSubmitRequest() *TaobaoSungariDisposeSubmitRequest{
    return &TaobaoSungariDisposeSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSungariDisposeSubmitRequest) GetApiMethodName() string {
    return "taobao.sungari.dispose.submit"
}

func (r TaobaoSungariDisposeSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSungariDisposeSubmitRequest) SetInfo(info *DisposeInfoDo) error {
    r.info = info
    r.Set("info", info)
    return nil
}

func (r TaobaoSungariDisposeSubmitRequest) GetInfo() *DisposeInfoDo {
    return r.info
}

