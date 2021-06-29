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
type TaobaoOauthCodeCreateRequest struct {
    model.Params
    // mock param
    test   int64
}

// 初始化TaobaoOauthCodeCreateRequest对象
func NewTaobaoOauthCodeCreateRequest() *TaobaoOauthCodeCreateRequest{
    return &TaobaoOauthCodeCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOauthCodeCreateRequest) GetApiMethodName() string {
    return "taobao.oauth.code.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOauthCodeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Test Setter
// mock param
func (r *TaobaoOauthCodeCreateRequest) SetTest(test int64) error {
    r.test = test
    r.Set("test", test)
    return nil
}

// Test Getter
func (r TaobaoOauthCodeCreateRequest) GetTest() int64 {
    return r.test
}
