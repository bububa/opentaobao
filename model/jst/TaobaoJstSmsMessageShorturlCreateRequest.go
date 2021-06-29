package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔营销效果短链生成 API请求
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

// 初始化TaobaoJstSmsMessageShorturlCreateRequest对象
func NewTaobaoJstSmsMessageShorturlCreateRequest() *TaobaoJstSmsMessageShorturlCreateRequest{
    return &TaobaoJstSmsMessageShorturlCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.shorturl.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NeedHttpsPrefix Setter
// 是否需要https前缀： true-要  false-不要
func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetNeedHttpsPrefix(needHttpsPrefix bool) error {
    r.needHttpsPrefix = needHttpsPrefix
    r.Set("need_https_prefix", needHttpsPrefix)
    return nil
}

// NeedHttpsPrefix Getter
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetNeedHttpsPrefix() bool {
    return r.needHttpsPrefix
}
// Tag Setter
// 人群标签
func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetTag(tag string) error {
    r.tag = tag
    r.Set("tag", tag)
    return nil
}

// Tag Getter
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetTag() string {
    return r.tag
}
// Url Setter
// 商品或者店铺的H5地址，只支持长链
func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetUrl() string {
    return r.url
}
// BatchNumber Setter
// 批次号
func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetBatchNumber(batchNumber string) error {
    r.batchNumber = batchNumber
    r.Set("batch_number", batchNumber)
    return nil
}

// BatchNumber Getter
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetBatchNumber() string {
    return r.batchNumber
}
