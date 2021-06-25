package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物联网绑定/换绑API APIRequest
alibaba.aliqin.fc.iot.modbind

支持用户的设备的换绑和解绑操作
*/
type AlibabaAliqinFcIotModbindRequest struct {
    model.Params

    // chgBind：换绑；unBind：解绑
    opionType   string 

    // 物联卡和iccid保持一致
    billReal   string 

    // iccid （20位）
    iccid   string 

    // 换绑的时候必传，换的新设备imei
    newimei   string 

    // 目前绑定的设备imei
    imei   string 

    // 物联卡业务：若无特殊为ICCID
    billSource   string 

    // 若无特殊物联卡传入122
    midPatChannel   string 

}

func NewAlibabaAliqinFcIotModbindRequest() *AlibabaAliqinFcIotModbindRequest{
    return &AlibabaAliqinFcIotModbindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcIotModbindRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.modbind"
}

func (r AlibabaAliqinFcIotModbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcIotModbindRequest) SetOpionType(opionType string) error {
    r.opionType = opionType
    r.Set("opion_type", opionType)
    return nil
}

func (r AlibabaAliqinFcIotModbindRequest) GetOpionType() string {
    return r.opionType
}

func (r *AlibabaAliqinFcIotModbindRequest) SetBillReal(billReal string) error {
    r.billReal = billReal
    r.Set("bill_real", billReal)
    return nil
}

func (r AlibabaAliqinFcIotModbindRequest) GetBillReal() string {
    return r.billReal
}

func (r *AlibabaAliqinFcIotModbindRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

func (r AlibabaAliqinFcIotModbindRequest) GetIccid() string {
    return r.iccid
}

func (r *AlibabaAliqinFcIotModbindRequest) SetNewimei(newimei string) error {
    r.newimei = newimei
    r.Set("newimei", newimei)
    return nil
}

func (r AlibabaAliqinFcIotModbindRequest) GetNewimei() string {
    return r.newimei
}

func (r *AlibabaAliqinFcIotModbindRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

func (r AlibabaAliqinFcIotModbindRequest) GetImei() string {
    return r.imei
}

func (r *AlibabaAliqinFcIotModbindRequest) SetBillSource(billSource string) error {
    r.billSource = billSource
    r.Set("bill_source", billSource)
    return nil
}

func (r AlibabaAliqinFcIotModbindRequest) GetBillSource() string {
    return r.billSource
}

func (r *AlibabaAliqinFcIotModbindRequest) SetMidPatChannel(midPatChannel string) error {
    r.midPatChannel = midPatChannel
    r.Set("mid_pat_channel", midPatChannel)
    return nil
}

func (r AlibabaAliqinFcIotModbindRequest) GetMidPatChannel() string {
    return r.midPatChannel
}

