package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡状态查询 APIRequest
alibaba.aliqin.fc.iot.cardStatus

物联卡状态查询
*/
type AlibabaAliqinFcIotCardStatusRequest struct {
    model.Params

    // SIM卡号
    iccid   string 

}

func NewAlibabaAliqinFcIotCardStatusRequest() *AlibabaAliqinFcIotCardStatusRequest{
    return &AlibabaAliqinFcIotCardStatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcIotCardStatusRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.cardStatus"
}

func (r AlibabaAliqinFcIotCardStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcIotCardStatusRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

func (r AlibabaAliqinFcIotCardStatusRequest) GetIccid() string {
    return r.iccid
}

