package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店信息接口 API请求
alibaba.ele.fengniao.chainstore.update

修改门店的经纬度，文本地址，电话，门店名
*/
type AlibabaEleFengniaoChainstoreUpdateRequest struct {
    model.Params
    // 入参
    param   *Param
}

// 初始化AlibabaEleFengniaoChainstoreUpdateRequest对象
func NewAlibabaEleFengniaoChainstoreUpdateRequest() *AlibabaEleFengniaoChainstoreUpdateRequest{
    return &AlibabaEleFengniaoChainstoreUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoChainstoreUpdateRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.chainstore.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoChainstoreUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaEleFengniaoChainstoreUpdateRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoChainstoreUpdateRequest) GetParam() *Param {
    return r.param
}
