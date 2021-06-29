package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
allspark提供抽奖tida接口对应鉴权接口 API请求
alibaba.interact.allsparkisv.draw

该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
*/
type AlibabaInteractAllsparkisvDrawRequest struct {
    model.Params
    // ddd
    _test   string
    // dd
    _ddd   string
}

// 初始化AlibabaInteractAllsparkisvDrawRequest对象
func NewAlibabaInteractAllsparkisvDrawRequest() *AlibabaInteractAllsparkisvDrawRequest{
    return &AlibabaInteractAllsparkisvDrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractAllsparkisvDrawRequest) GetApiMethodName() string {
    return "alibaba.interact.allsparkisv.draw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractAllsparkisvDrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Test Setter
// ddd
func (r *AlibabaInteractAllsparkisvDrawRequest) SetTest(_test string) error {
    r._test = _test
    r.Set("test", _test)
    return nil
}

// Test Getter
func (r AlibabaInteractAllsparkisvDrawRequest) GetTest() string {
    return r._test
}
// Ddd Setter
// dd
func (r *AlibabaInteractAllsparkisvDrawRequest) SetDdd(_ddd string) error {
    r._ddd = _ddd
    r.Set("ddd", _ddd)
    return nil
}

// Ddd Getter
func (r AlibabaInteractAllsparkisvDrawRequest) GetDdd() string {
    return r._ddd
}
