package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医院药品子码申请接口 API请求
alibaba.alihealth.drug.code.kyt.yy.applycode

根据父码及所属企业ID生成子码信息
*/
type AlibabaAlihealthDrugCodeKytYyApplycodeRequest struct {
    model.Params
    // 企业ID（ref_ent_id)
    _refEntId   string
    // 父码
    _code   string
    // 申请数量
    _amount   int64
}

// 初始化AlibabaAlihealthDrugCodeKytYyApplycodeRequest对象
func NewAlibabaAlihealthDrugCodeKytYyApplycodeRequest() *AlibabaAlihealthDrugCodeKytYyApplycodeRequest{
    return &AlibabaAlihealthDrugCodeKytYyApplycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.yy.applycode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID（ref_ent_id)
func (r *AlibabaAlihealthDrugCodeKytYyApplycodeRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetRefEntId() string {
    return r._refEntId
}
// Code Setter
// 父码
func (r *AlibabaAlihealthDrugCodeKytYyApplycodeRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetCode() string {
    return r._code
}
// Amount Setter
// 申请数量
func (r *AlibabaAlihealthDrugCodeKytYyApplycodeRequest) SetAmount(_amount int64) error {
    r._amount = _amount
    r.Set("amount", _amount)
    return nil
}

// Amount Getter
func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetAmount() int64 {
    return r._amount
}
