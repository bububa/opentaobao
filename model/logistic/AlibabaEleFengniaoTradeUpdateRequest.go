package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新蜂鸟扣费状态 APIRequest
alibaba.ele.fengniao.trade.update

汇金扣费成功后，回调该接口更新扣费状态
*/
type AlibabaEleFengniaoTradeUpdateRequest struct {
    model.Params

    // param 参数
    param   *Param 

}

func NewAlibabaEleFengniaoTradeUpdateRequest() *AlibabaEleFengniaoTradeUpdateRequest{
    return &AlibabaEleFengniaoTradeUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoTradeUpdateRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.trade.update"
}

func (r AlibabaEleFengniaoTradeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoTradeUpdateRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoTradeUpdateRequest) GetParam() *Param {
    return r.param
}

