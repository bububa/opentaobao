package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中心化供应商异常号码状态同步接口 APIRequest
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

func NewAlibabaAliqinAxbVendorExceptionNoSyncRequest() *AlibabaAliqinAxbVendorExceptionNoSyncRequest{
    return &AlibabaAliqinAxbVendorExceptionNoSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.exception.no.sync"
}

func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetSecretNo(secretNo string) error {
    r.secretNo = secretNo
    r.Set("secret_no", secretNo)
    return nil
}

func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetSecretNo() string {
    return r.secretNo
}

func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetExceptionMsg(exceptionMsg string) error {
    r.exceptionMsg = exceptionMsg
    r.Set("exception_msg", exceptionMsg)
    return nil
}

func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetExceptionMsg() string {
    return r.exceptionMsg
}

func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaAliqinAxbVendorExceptionNoSyncRequest) SetVendorKey(vendorKey string) error {
    r.vendorKey = vendorKey
    r.Set("vendor_key", vendorKey)
    return nil
}

func (r AlibabaAliqinAxbVendorExceptionNoSyncRequest) GetVendorKey() string {
    return r.vendorKey
}

