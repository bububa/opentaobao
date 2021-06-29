package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预采购服务包查询接口 API请求
alibaba.ele.fengniao.service.package.query

查询门店所在经纬度可用服务包的接口
*/
type AlibabaEleFengniaoServicePackageQueryRequest struct {
    model.Params
    // 入参
    param   *Param
}

// 初始化AlibabaEleFengniaoServicePackageQueryRequest对象
func NewAlibabaEleFengniaoServicePackageQueryRequest() *AlibabaEleFengniaoServicePackageQueryRequest{
    return &AlibabaEleFengniaoServicePackageQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoServicePackageQueryRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.service.package.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoServicePackageQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaEleFengniaoServicePackageQueryRequest) SetParam(param *Param) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaEleFengniaoServicePackageQueryRequest) GetParam() *Param {
    return r.param
}
