package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索接口（首字母） APIRequest
alibaba.xiami.api.search.letter.get

搜索接口（首字母）
*/
type AlibabaXiamiApiSearchLetterGetRequest struct {
    model.Params

    // 搜索关键字
    key   string 

}

func NewAlibabaXiamiApiSearchLetterGetRequest() *AlibabaXiamiApiSearchLetterGetRequest{
    return &AlibabaXiamiApiSearchLetterGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaXiamiApiSearchLetterGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.search.letter.get"
}

func (r AlibabaXiamiApiSearchLetterGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaXiamiApiSearchLetterGetRequest) SetKey(key string) error {
    r.key = key
    r.Set("key", key)
    return nil
}

func (r AlibabaXiamiApiSearchLetterGetRequest) GetKey() string {
    return r.key
}

