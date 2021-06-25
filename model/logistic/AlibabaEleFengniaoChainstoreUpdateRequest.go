package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店信息接口 APIRequest
alibaba.ele.fengniao.chainstore.update

修改门店的经纬度，文本地址，电话，门店名
*/
type AlibabaEleFengniaoChainstoreUpdateRequest struct {
    model.Params

    // 入参
    param   *Param 

}

func NewAlibabaEleFengniaoChainstoreUpdateRequest() *AlibabaEleFengniaoChainstoreUpdateRequest{
    return &AlibabaEleFengniaoChainstoreUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoChainstoreUpdateRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.chainstore.update"
}

func (r AlibabaEleFengniaoChainstoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoChainstoreUpdateRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaEleFengniaoChainstoreUpdateRequest) GetParam() *Param {
    return r.param
}

