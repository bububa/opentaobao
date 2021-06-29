package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物联网绑定/换绑API API请求
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

// 初始化AlibabaAliqinFcIotModbindRequest对象
func NewAlibabaAliqinFcIotModbindRequest() *AlibabaAliqinFcIotModbindRequest{
    return &AlibabaAliqinFcIotModbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotModbindRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.modbind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotModbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpionType Setter
// chgBind：换绑；unBind：解绑
func (r *AlibabaAliqinFcIotModbindRequest) SetOpionType(opionType string) error {
    r.opionType = opionType
    r.Set("opion_type", opionType)
    return nil
}

// OpionType Getter
func (r AlibabaAliqinFcIotModbindRequest) GetOpionType() string {
    return r.opionType
}
// BillReal Setter
// 物联卡和iccid保持一致
func (r *AlibabaAliqinFcIotModbindRequest) SetBillReal(billReal string) error {
    r.billReal = billReal
    r.Set("bill_real", billReal)
    return nil
}

// BillReal Getter
func (r AlibabaAliqinFcIotModbindRequest) GetBillReal() string {
    return r.billReal
}
// Iccid Setter
// iccid （20位）
func (r *AlibabaAliqinFcIotModbindRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotModbindRequest) GetIccid() string {
    return r.iccid
}
// Newimei Setter
// 换绑的时候必传，换的新设备imei
func (r *AlibabaAliqinFcIotModbindRequest) SetNewimei(newimei string) error {
    r.newimei = newimei
    r.Set("newimei", newimei)
    return nil
}

// Newimei Getter
func (r AlibabaAliqinFcIotModbindRequest) GetNewimei() string {
    return r.newimei
}
// Imei Setter
// 目前绑定的设备imei
func (r *AlibabaAliqinFcIotModbindRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

// Imei Getter
func (r AlibabaAliqinFcIotModbindRequest) GetImei() string {
    return r.imei
}
// BillSource Setter
// 物联卡业务：若无特殊为ICCID
func (r *AlibabaAliqinFcIotModbindRequest) SetBillSource(billSource string) error {
    r.billSource = billSource
    r.Set("bill_source", billSource)
    return nil
}

// BillSource Getter
func (r AlibabaAliqinFcIotModbindRequest) GetBillSource() string {
    return r.billSource
}
// MidPatChannel Setter
// 若无特殊物联卡传入122
func (r *AlibabaAliqinFcIotModbindRequest) SetMidPatChannel(midPatChannel string) error {
    r.midPatChannel = midPatChannel
    r.Set("mid_pat_channel", midPatChannel)
    return nil
}

// MidPatChannel Getter
func (r AlibabaAliqinFcIotModbindRequest) GetMidPatChannel() string {
    return r.midPatChannel
}
