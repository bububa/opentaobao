package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡信息查询 APIRequest
alibaba.aliqin.fc.iot.cardInfo

物联卡信息查询
*/
type AlibabaAliqinFcIotCardInfoRequest struct {
    model.Params

    // SIM卡号
    iccid   string 

}

func NewAlibabaAliqinFcIotCardInfoRequest() *AlibabaAliqinFcIotCardInfoRequest{
    return &AlibabaAliqinFcIotCardInfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcIotCardInfoRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.cardInfo"
}

func (r AlibabaAliqinFcIotCardInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcIotCardInfoRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

func (r AlibabaAliqinFcIotCardInfoRequest) GetIccid() string {
    return r.iccid
}

