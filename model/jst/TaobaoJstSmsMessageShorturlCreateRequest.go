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
    _needHttpsPrefix   bool
    // 人群标签
    _tag   string
    // 商品或者店铺的H5地址，只支持长链
    _url   string
    // 批次号
    _batchNumber   string
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
func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetNeedHttpsPrefix(_needHttpsPrefix bool) error {
    r._needHttpsPrefix = _needHttpsPrefix
    r.Set("need_https_prefix", _needHttpsPrefix)
    return nil
}

// NeedHttpsPrefix Getter
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetNeedHttpsPrefix() bool {
    return r._needHttpsPrefix
}
// Tag Setter
// 人群标签
func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetTag(_tag string) error {
    r._tag = _tag
    r.Set("tag", _tag)
    return nil
}

// Tag Getter
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetTag() string {
    return r._tag
}
// Url Setter
// 商品或者店铺的H5地址，只支持长链
func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetUrl() string {
    return r._url
}
// BatchNumber Setter
// 批次号
func (r *TaobaoJstSmsMessageShorturlCreateRequest) SetBatchNumber(_batchNumber string) error {
    r._batchNumber = _batchNumber
    r.Set("batch_number", _batchNumber)
    return nil
}

// BatchNumber Getter
func (r TaobaoJstSmsMessageShorturlCreateRequest) GetBatchNumber() string {
    return r._batchNumber
}
