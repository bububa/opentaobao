package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中心化供应商异常号码状态同步接口 API请求
alibaba.aliqin.axb.vendor.exception.no.sync

用于中心化供应商同步异常号码
*/
type AlibabaAliqinAxbVendorExceptionNoSyncRequest struct {
    model.Params
    // 异常的中间号码
    _secretNo   string
    // 异常的原因
    _exceptionMsg   string
    // 0-异常状态 1-可恢复正常使用
    _status   int64
    // 供应商KEY
    _vendorKey   string
}

// 初始化AlibabaAliqinAxbVendorExceptionNoSyncRequest对象
func NewAlibabaAliqinAxbVendorExceptionNoSyncRequest() *AlibabaAliqinAxbVendorExceptionNoSyncRequest{
    return &AlibabaAliqinAxbVendorExceptionNoSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.exception.no.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SecretNo Setter
// 异常的中间号码
func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetSecretNo(_secretNo string) error {
    r._secretNo = _secretNo
    r.Set("secret_no", _secretNo)
    return nil
}

// SecretNo Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetSecretNo() string {
    return r._secretNo
}
// ExceptionMsg Setter
// 异常的原因
func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetExceptionMsg(_exceptionMsg string) error {
    r._exceptionMsg = _exceptionMsg
    r.Set("exception_msg", _exceptionMsg)
    return nil
}

// ExceptionMsg Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetExceptionMsg() string {
    return r._exceptionMsg
}
// Status Setter
// 0-异常状态 1-可恢复正常使用
func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetStatus() int64 {
    return r._status
}
// VendorKey Setter
// 供应商KEY
func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetVendorKey(_vendorKey string) error {
    r._vendorKey = _vendorKey
    r.Set("vendor_key", _vendorKey)
    return nil
}

// VendorKey Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetVendorKey() string {
    return r._vendorKey
}
