package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询终端信息 APIRequest
alibaba.aliqin.fc.iot.qrycard

查询终端信息
*/
type AlibabaAliqinFcIotQrycardRequest struct {
    model.Params

    // 外部计费来源
    billSource   string 

    // 外部计费号
    billReal   string 

    // ICCID
    iccid   string 

}

func NewAlibabaAliqinFcIotQrycardRequest() *AlibabaAliqinFcIotQrycardRequest{
    return &AlibabaAliqinFcIotQrycardRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcIotQrycardRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.qrycard"
}

func (r AlibabaAliqinFcIotQrycardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcIotQrycardRequest) SetBillSource(billSource string) error {
    r.billSource = billSource
    r.Set("bill_source", billSource)
    return nil
}

func (r AlibabaAliqinFcIotQrycardRequest) GetBillSource() string {
    return r.billSource
}

func (r *AlibabaAliqinFcIotQrycardRequest) SetBillReal(billReal string) error {
    r.billReal = billReal
    r.Set("bill_real", billReal)
    return nil
}

func (r AlibabaAliqinFcIotQrycardRequest) GetBillReal() string {
    return r.billReal
}

func (r *AlibabaAliqinFcIotQrycardRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

func (r AlibabaAliqinFcIotQrycardRequest) GetIccid() string {
    return r.iccid
}

