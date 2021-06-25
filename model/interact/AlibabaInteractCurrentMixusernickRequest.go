package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘混淆nick开放接口鉴权专用 APIRequest
alibaba.interact.current.mixusernick

手淘混淆nick开放接口鉴权专用，无数据输入输出。
*/
type AlibabaInteractCurrentMixusernickRequest struct {
    model.Params

}

func NewAlibabaInteractCurrentMixusernickRequest() *AlibabaInteractCurrentMixusernickRequest{
    return &AlibabaInteractCurrentMixusernickRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractCurrentMixusernickRequest) GetApiMethodName() string {
    return "alibaba.interact.current.mixusernick"
}

func (r AlibabaInteractCurrentMixusernickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


