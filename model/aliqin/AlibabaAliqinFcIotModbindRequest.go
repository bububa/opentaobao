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
    _opionType   string
    // 物联卡和iccid保持一致
    _billReal   string
    // iccid （20位）
    _iccid   string
    // 换绑的时候必传，换的新设备imei
    _newimei   string
    // 目前绑定的设备imei
    _imei   string
    // 物联卡业务：若无特殊为ICCID
    _billSource   string
    // 若无特殊物联卡传入122
    _midPatChannel   string
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
func (r *AlibabaAliqinFcIotModbindRequest) SetOpionType(_opionType string) error {
    r._opionType = _opionType
    r.Set("opion_type", _opionType)
    return nil
}

// OpionType Getter
func (r AlibabaAliqinFcIotModbindRequest) GetOpionType() string {
    return r._opionType
}
// BillReal Setter
// 物联卡和iccid保持一致
func (r *AlibabaAliqinFcIotModbindRequest) SetBillReal(_billReal string) error {
    r._billReal = _billReal
    r.Set("bill_real", _billReal)
    return nil
}

// BillReal Getter
func (r AlibabaAliqinFcIotModbindRequest) GetBillReal() string {
    return r._billReal
}
// Iccid Setter
// iccid （20位）
func (r *AlibabaAliqinFcIotModbindRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotModbindRequest) GetIccid() string {
    return r._iccid
}
// Newimei Setter
// 换绑的时候必传，换的新设备imei
func (r *AlibabaAliqinFcIotModbindRequest) SetNewimei(_newimei string) error {
    r._newimei = _newimei
    r.Set("newimei", _newimei)
    return nil
}

// Newimei Getter
func (r AlibabaAliqinFcIotModbindRequest) GetNewimei() string {
    return r._newimei
}
// Imei Setter
// 目前绑定的设备imei
func (r *AlibabaAliqinFcIotModbindRequest) SetImei(_imei string) error {
    r._imei = _imei
    r.Set("imei", _imei)
    return nil
}

// Imei Getter
func (r AlibabaAliqinFcIotModbindRequest) GetImei() string {
    return r._imei
}
// BillSource Setter
// 物联卡业务：若无特殊为ICCID
func (r *AlibabaAliqinFcIotModbindRequest) SetBillSource(_billSource string) error {
    r._billSource = _billSource
    r.Set("bill_source", _billSource)
    return nil
}

// BillSource Getter
func (r AlibabaAliqinFcIotModbindRequest) GetBillSource() string {
    return r._billSource
}
// MidPatChannel Setter
// 若无特殊物联卡传入122
func (r *AlibabaAliqinFcIotModbindRequest) SetMidPatChannel(_midPatChannel string) error {
    r._midPatChannel = _midPatChannel
    r.Set("mid_pat_channel", _midPatChannel)
    return nil
}

// MidPatChannel Getter
func (r AlibabaAliqinFcIotModbindRequest) GetMidPatChannel() string {
    return r._midPatChannel
}
