package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询终端信息 API请求
alibaba.aliqin.fc.iot.qrycard

查询终端信息
*/
type AlibabaAliqinFcIotQrycardRequest struct {
    model.Params
    // 外部计费来源
    _billSource   string
    // 外部计费号
    _billReal   string
    // ICCID
    _iccid   string
}

// 初始化AlibabaAliqinFcIotQrycardRequest对象
func NewAlibabaAliqinFcIotQrycardRequest() *AlibabaAliqinFcIotQrycardRequest{
    return &AlibabaAliqinFcIotQrycardRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotQrycardRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.qrycard"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotQrycardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillSource Setter
// 外部计费来源
func (r *AlibabaAliqinFcIotQrycardRequest) SetBillSource(_billSource string) error {
    r._billSource = _billSource
    r.Set("bill_source", _billSource)
    return nil
}

// BillSource Getter
func (r AlibabaAliqinFcIotQrycardRequest) GetBillSource() string {
    return r._billSource
}
// BillReal Setter
// 外部计费号
func (r *AlibabaAliqinFcIotQrycardRequest) SetBillReal(_billReal string) error {
    r._billReal = _billReal
    r.Set("bill_real", _billReal)
    return nil
}

// BillReal Getter
func (r AlibabaAliqinFcIotQrycardRequest) GetBillReal() string {
    return r._billReal
}
// Iccid Setter
// ICCID
func (r *AlibabaAliqinFcIotQrycardRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotQrycardRequest) GetIccid() string {
    return r._iccid
}
