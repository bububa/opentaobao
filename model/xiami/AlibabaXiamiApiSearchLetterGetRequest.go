package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索接口（首字母） API请求
alibaba.xiami.api.search.letter.get

搜索接口（首字母）
*/
type AlibabaXiamiApiSearchLetterGetRequest struct {
    model.Params
    // 搜索关键字
    key   string
}

// 初始化AlibabaXiamiApiSearchLetterGetRequest对象
func NewAlibabaXiamiApiSearchLetterGetRequest() *AlibabaXiamiApiSearchLetterGetRequest{
    return &AlibabaXiamiApiSearchLetterGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiSearchLetterGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.search.letter.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiSearchLetterGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Key Setter
// 搜索关键字
func (r *AlibabaXiamiApiSearchLetterGetRequest) SetKey(key string) error {
    r.key = key
    r.Set("key", key)
    return nil
}

// Key Getter
func (r AlibabaXiamiApiSearchLetterGetRequest) GetKey() string {
    return r.key
}
