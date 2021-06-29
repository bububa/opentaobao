package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短链接获取 API请求
alibaba.retail.shorturl.get

短链接获取
*/
type AlibabaRetailShorturlGetRequest struct {
    model.Params
    // 源url
    _sourceUrl   string
    // 系统自动生成
    _options   *ShortUrlOption
}

// 初始化AlibabaRetailShorturlGetRequest对象
func NewAlibabaRetailShorturlGetRequest() *AlibabaRetailShorturlGetRequest{
    return &AlibabaRetailShorturlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailShorturlGetRequest) GetApiMethodName() string {
    return "alibaba.retail.shorturl.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailShorturlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SourceUrl Setter
// 源url
func (r *AlibabaRetailShorturlGetRequest) SetSourceUrl(_sourceUrl string) error {
    r._sourceUrl = _sourceUrl
    r.Set("source_url", _sourceUrl)
    return nil
}

// SourceUrl Getter
func (r AlibabaRetailShorturlGetRequest) GetSourceUrl() string {
    return r._sourceUrl
}
// Options Setter
// 系统自动生成
func (r *AlibabaRetailShorturlGetRequest) SetOptions(_options *ShortUrlOption) error {
    r._options = _options
    r.Set("options", _options)
    return nil
}

// Options Getter
func (r AlibabaRetailShorturlGetRequest) GetOptions() *ShortUrlOption {
    return r._options
}
