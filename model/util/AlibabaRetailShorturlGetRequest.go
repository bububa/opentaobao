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
    sourceUrl   string
    // 系统自动生成
    options   *ShortUrlOption
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
func (r *AlibabaRetailShorturlGetRequest) SetSourceUrl(sourceUrl string) error {
    r.sourceUrl = sourceUrl
    r.Set("source_url", sourceUrl)
    return nil
}

// SourceUrl Getter
func (r AlibabaRetailShorturlGetRequest) GetSourceUrl() string {
    return r.sourceUrl
}
// Options Setter
// 系统自动生成
func (r *AlibabaRetailShorturlGetRequest) SetOptions(options *ShortUrlOption) error {
    r.options = options
    r.Set("options", options)
    return nil
}

// Options Getter
func (r AlibabaRetailShorturlGetRequest) GetOptions() *ShortUrlOption {
    return r.options
}
