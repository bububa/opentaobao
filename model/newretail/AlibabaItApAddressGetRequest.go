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
    signature   string
    // 语言
    language   string
    // 分配的ak
    appKeyInternal   string
    // ap的mac
    mac   string
    // 当前时间戳，毫秒
    timestampInternal   int64
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
func (r *AlibabaItApAddressGetRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

// Signature Getter
func (r AlibabaItApAddressGetRequest) GetSignature() string {
    return r.signature
}
// Language Setter
// 语言
func (r *AlibabaItApAddressGetRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaItApAddressGetRequest) GetLanguage() string {
    return r.language
}
// AppKeyInternal Setter
// 分配的ak
func (r *AlibabaItApAddressGetRequest) SetAppKeyInternal(appKeyInternal string) error {
    r.appKeyInternal = appKeyInternal
    r.Set("app_key_internal", appKeyInternal)
    return nil
}

// AppKeyInternal Getter
func (r AlibabaItApAddressGetRequest) GetAppKeyInternal() string {
    return r.appKeyInternal
}
// Mac Setter
// ap的mac
func (r *AlibabaItApAddressGetRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

// Mac Getter
func (r AlibabaItApAddressGetRequest) GetMac() string {
    return r.mac
}
// TimestampInternal Setter
// 当前时间戳，毫秒
func (r *AlibabaItApAddressGetRequest) SetTimestampInternal(timestampInternal int64) error {
    r.timestampInternal = timestampInternal
    r.Set("timestamp_internal", timestampInternal)
    return nil
}

// TimestampInternal Getter
func (r AlibabaItApAddressGetRequest) GetTimestampInternal() int64 {
    return r.timestampInternal
}
