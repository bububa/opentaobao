package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物联网卡上订购的offer APIRequest
alibaba.aliqin.fc.iot.cardoffer

查询物联网卡上订购的offer
*/
type AlibabaAliqinFcIotCardofferRequest struct {
    model.Params

    // 具体ICCID的值
    billreal   string 

    // ICCID
    billsource   string 

}

func NewAlibabaAliqinFcIotCardofferRequest() *AlibabaAliqinFcIotCardofferRequest{
    return &AlibabaAliqinFcIotCardofferRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcIotCardofferRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.cardoffer"
}

func (r AlibabaAliqinFcIotCardofferRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcIotCardofferRequest) SetBillreal(billreal string) error {
    r.billreal = billreal
    r.Set("billreal", billreal)
    return nil
}

func (r AlibabaAliqinFcIotCardofferRequest) GetBillreal() string {
    return r.billreal
}

func (r *AlibabaAliqinFcIotCardofferRequest) SetBillsource(billsource string) error {
    r.billsource = billsource
    r.Set("billsource", billsource)
    return nil
}

func (r AlibabaAliqinFcIotCardofferRequest) GetBillsource() string {
    return r.billsource
}

