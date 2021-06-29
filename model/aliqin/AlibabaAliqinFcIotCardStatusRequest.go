package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡状态查询 API请求
alibaba.aliqin.fc.iot.cardStatus

物联卡状态查询
*/
type AlibabaAliqinFcIotCardStatusRequest struct {
    model.Params
    // SIM卡号
    _iccid   string
}

// 初始化AlibabaAliqinFcIotCardStatusRequest对象
func NewAlibabaAliqinFcIotCardStatusRequest() *AlibabaAliqinFcIotCardStatusRequest{
    return &AlibabaAliqinFcIotCardStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotCardStatusRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.cardStatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotCardStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Iccid Setter
// SIM卡号
func (r *AlibabaAliqinFcIotCardStatusRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotCardStatusRequest) GetIccid() string {
    return r._iccid
}
