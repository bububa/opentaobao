package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔营销效果短链生成 APIRequest
taobao.jst.sms.message.shorturl.create

聚石塔生成淘短链接口
*/
type TaobaoJstSmsMessageShorturlCreateRequest struct {
    model.Params

    // 是否需要https前缀： true-要  false-不要
    needHttpsPrefix   bool 

    // 人群标签
    tag   string 

    // 商品或者店铺的H5地址，只支持长链
    url   string 

    // 批次号
    batchNumber   string 

}

func NewTaobaoJstSmsMessageShorturlCreateRequest() *TaobaoJstSmsMessageShorturlCreateRequest{
    return &TaobaoJstSmsMessageShorturlCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsMessageShorturlCreateRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.shorturl.create"
}

func (r TaobaoJstSmsMessageShorturlCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetNeedHttpsPrefix(needHttpsPrefix bool) error {
    r.needHttpsPrefix = needHttpsPrefix
    r.Set("need_https_prefix", needHttpsPrefix)
    return nil
}

func (r TaobaoJstSmsMessageShorturlCreateRequest) GetNeedHttpsPrefix() bool {
    return r.needHttpsPrefix
}

func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetTag(tag string) error {
    r.tag = tag
    r.Set("tag", tag)
    return nil
}

func (r TaobaoJstSmsMessageShorturlCreateRequest) GetTag() string {
    return r.tag
}

func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TaobaoJstSmsMessageShorturlCreateRequest) GetUrl() string {
    return r.url
}

func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetBatchNumber(batchNumber string) error {
    r.batchNumber = batchNumber
    r.Set("batch_number", batchNumber)
    return nil
}

func (r TaobaoJstSmsMessageShorturlCreateRequest) GetBatchNumber() string {
    return r.batchNumber
}

