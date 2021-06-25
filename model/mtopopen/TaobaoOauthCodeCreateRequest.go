package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝OauthCode颁发 APIRequest
taobao.oauth.code.create

手淘无线开放的oauthCode颁发接口
*/
type TaobaoOauthCodeCreateRequest struct {
    model.Params

    // mock param
    test   int64 

}

func NewTaobaoOauthCodeCreateRequest() *TaobaoOauthCodeCreateRequest{
    return &TaobaoOauthCodeCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOauthCodeCreateRequest) GetApiMethodName() string {
    return "taobao.oauth.code.create"
}

func (r TaobaoOauthCodeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOauthCodeCreateRequest) SetTest(test int64) error {
    r.test = test
    r.Set("test", test)
    return nil
}

func (r TaobaoOauthCodeCreateRequest) GetTest() int64 {
    return r.test
}

