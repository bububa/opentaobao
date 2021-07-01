package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝OauthCode颁发 API请求
taobao.oauth.code.create

手淘无线开放的oauthCode颁发接口
*/
type TaobaoOauthCodeCreateAPIRequest struct {
    model.Params
    // mock param
    _test   int64
}

// 初始化TaobaoOauthCodeCreateAPIRequest对象
func NewTaobaoOauthCodeCreateRequest() *TaobaoOauthCodeCreateAPIRequest{
    return &TaobaoOauthCodeCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOauthCodeCreateAPIRequest) GetApiMethodName() string {
    return "taobao.oauth.code.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOauthCodeCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Test Setter
// mock param
func (r *TaobaoOauthCodeCreateAPIRequest) SetTest(_test int64) error {
    r._test = _test
    r.Set("test", _test)
    return nil
}

// Test Getter
func (r TaobaoOauthCodeCreateAPIRequest) GetTest() int64 {
    return r._test
}
