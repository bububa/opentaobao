package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡信息查询 API请求
alibaba.aliqin.fc.iot.cardInfo

物联卡信息查询
*/
type AlibabaAliqinFcIotCardInfoRequest struct {
    model.Params
    // SIM卡号
    _iccid   string
}

// 初始化AlibabaAliqinFcIotCardInfoRequest对象
func NewAlibabaAliqinFcIotCardInfoRequest() *AlibabaAliqinFcIotCardInfoRequest{
    return &AlibabaAliqinFcIotCardInfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotCardInfoRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.cardInfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotCardInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Iccid Setter
// SIM卡号
func (r *AlibabaAliqinFcIotCardInfoRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotCardInfoRequest) GetIccid() string {
    return r._iccid
}
