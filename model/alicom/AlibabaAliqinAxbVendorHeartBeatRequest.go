package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商心跳上报接口 APIRequest
alibaba.aliqin.axb.vendor.heart.beat

供应商上报自己的心跳信息
*/
type AlibabaAliqinAxbVendorHeartBeatRequest struct {
    model.Params

    // 可选的预留字段
    status   string 

    // 供应商合作KEY
    vendorKey   string 

}

func NewAlibabaAliqinAxbVendorHeartBeatRequest() *AlibabaAliqinAxbVendorHeartBeatRequest{
    return &AlibabaAliqinAxbVendorHeartBeatRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinAxbVendorHeartBeatRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.heart.beat"
}

func (r AlibabaAliqinAxbVendorHeartBeatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinAxbVendorHeartBeatRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaAliqinAxbVendorHeartBeatRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaAliqinAxbVendorHeartBeatRequest) SetVendorKey(vendorKey string) error {
    r.vendorKey = vendorKey
    r.Set("vendor_key", vendorKey)
    return nil
}

func (r AlibabaAliqinAxbVendorHeartBeatRequest) GetVendorKey() string {
    return r.vendorKey
}

