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
    secretNo   string
    // 异常的原因
    exceptionMsg   string
    // 0-异常状态 1-可恢复正常使用
    status   int64
    // 供应商KEY
    vendorKey   string
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
func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetSecretNo(secretNo string) error {
    r.secretNo = secretNo
    r.Set("secret_no", secretNo)
    return nil
}

// SecretNo Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetSecretNo() string {
    return r.secretNo
}
// ExceptionMsg Setter
// 异常的原因
func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetExceptionMsg(exceptionMsg string) error {
    r.exceptionMsg = exceptionMsg
    r.Set("exception_msg", exceptionMsg)
    return nil
}

// ExceptionMsg Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetExceptionMsg() string {
    return r.exceptionMsg
}
// Status Setter
// 0-异常状态 1-可恢复正常使用
func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetStatus() int64 {
    return r.status
}
// VendorKey Setter
// 供应商KEY
func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetVendorKey(vendorKey string) error {
    r.vendorKey = vendorKey
    r.Set("vendor_key", vendorKey)
    return nil
}

// VendorKey Getter
func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetVendorKey() string {
    return r.vendorKey
}
