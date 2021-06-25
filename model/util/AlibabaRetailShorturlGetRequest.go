package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
短链接获取 APIRequest
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

func NewAlibabaRetailShorturlGetRequest() *AlibabaRetailShorturlGetRequest{
    return &AlibabaRetailShorturlGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailShorturlGetRequest) GetApiMethodName() string {
    return "alibaba.retail.shorturl.get"
}

func (r AlibabaRetailShorturlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailShorturlGetRequest) SetSourceUrl(sourceUrl string) error {
    r.sourceUrl = sourceUrl
    r.Set("source_url", sourceUrl)
    return nil
}

func (r AlibabaRetailShorturlGetRequest) GetSourceUrl() string {
    return r.sourceUrl
}

func (r *AlibabaRetailShorturlGetRequest) SetOptions(options *ShortUrlOption) error {
    r.options = options
    r.Set("options", options)
    return nil
}

func (r AlibabaRetailShorturlGetRequest) GetOptions() *ShortUrlOption {
    return r.options
}

