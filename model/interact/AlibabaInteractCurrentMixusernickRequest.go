package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘混淆nick开放接口鉴权专用 API请求
alibaba.interact.current.mixusernick

手淘混淆nick开放接口鉴权专用，无数据输入输出。
*/
type AlibabaInteractCurrentMixusernickRequest struct {
    model.Params
}

// 初始化AlibabaInteractCurrentMixusernickRequest对象
func NewAlibabaInteractCurrentMixusernickRequest() *AlibabaInteractCurrentMixusernickRequest{
    return &AlibabaInteractCurrentMixusernickRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractCurrentMixusernickRequest) GetApiMethodName() string {
    return "alibaba.interact.current.mixusernick"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractCurrentMixusernickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
