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
    billSource   string
    // 外部计费号
    billReal   string
    // ICCID
    iccid   string
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
func (r *AlibabaAliqinFcIotQrycardRequest) SetBillSource(billSource string) error {
    r.billSource = billSource
    r.Set("bill_source", billSource)
    return nil
}

// BillSource Getter
func (r AlibabaAliqinFcIotQrycardRequest) GetBillSource() string {
    return r.billSource
}
// BillReal Setter
// 外部计费号
func (r *AlibabaAliqinFcIotQrycardRequest) SetBillReal(billReal string) error {
    r.billReal = billReal
    r.Set("bill_real", billReal)
    return nil
}

// BillReal Getter
func (r AlibabaAliqinFcIotQrycardRequest) GetBillReal() string {
    return r.billReal
}
// Iccid Setter
// ICCID
func (r *AlibabaAliqinFcIotQrycardRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotQrycardRequest) GetIccid() string {
    return r.iccid
}
