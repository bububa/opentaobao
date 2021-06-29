package newretail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
getApAddressByMacNew API请求
alibaba.it.ap.address.get

根据ap 的mac地址查询ap的结构化位置信息
*/
type AlibabaItApAddressGetRequest struct {
    model.Params
    // 签名
    _signature   string
    // 语言
    _language   string
    // 分配的ak
    _appKeyInternal   string
    // ap的mac
    _mac   string
    // 当前时间戳，毫秒
    _timestampInternal   int64
}

// 初始化AlibabaItApAddressGetRequest对象
func NewAlibabaItApAddressGetRequest() *AlibabaItApAddressGetRequest{
    return &AlibabaItApAddressGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItApAddressGetRequest) GetApiMethodName() string {
    return "alibaba.it.ap.address.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItApAddressGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Signature Setter
// 签名
func (r *AlibabaItApAddressGetRequest) SetSignature(_signature string) error {
    r._signature = _signature
    r.Set("signature", _signature)
    return nil
}

// Signature Getter
func (r AlibabaItApAddressGetRequest) GetSignature() string {
    return r._signature
}
// Language Setter
// 语言
func (r *AlibabaItApAddressGetRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AlibabaItApAddressGetRequest) GetLanguage() string {
    return r._language
}
// AppKeyInternal Setter
// 分配的ak
func (r *AlibabaItApAddressGetRequest) SetAppKeyInternal(_appKeyInternal string) error {
    r._appKeyInternal = _appKeyInternal
    r.Set("app_key_internal", _appKeyInternal)
    return nil
}

// AppKeyInternal Getter
func (r AlibabaItApAddressGetRequest) GetAppKeyInternal() string {
    return r._appKeyInternal
}
// Mac Setter
// ap的mac
func (r *AlibabaItApAddressGetRequest) SetMac(_mac string) error {
    r._mac = _mac
    r.Set("mac", _mac)
    return nil
}

// Mac Getter
func (r AlibabaItApAddressGetRequest) GetMac() string {
    return r._mac
}
// TimestampInternal Setter
// 当前时间戳，毫秒
func (r *AlibabaItApAddressGetRequest) SetTimestampInternal(_timestampInternal int64) error {
    r._timestampInternal = _timestampInternal
    r.Set("timestamp_internal", _timestampInternal)
    return nil
}

// TimestampInternal Getter
func (r AlibabaItApAddressGetRequest) GetTimestampInternal() int64 {
    return r._timestampInternal
}
