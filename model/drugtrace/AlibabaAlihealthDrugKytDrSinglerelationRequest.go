package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多融单码关联关系查询 API请求
alibaba.alihealth.drug.kyt.dr.singlerelation

单码关联关系查询
*/
type AlibabaAlihealthDrugKytDrSinglerelationRequest struct {
    model.Params
    // 追溯码
    _code   string
    // 接口调用企业的唯一标识（接口调用者）
    _refEntId   string
    // 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
    _desRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytDrSinglerelationRequest对象
func NewAlibabaAlihealthDrugKytDrSinglerelationRequest() *AlibabaAlihealthDrugKytDrSinglerelationRequest{
    return &AlibabaAlihealthDrugKytDrSinglerelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.singlerelation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytDrSinglerelationRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetCode() string {
    return r._code
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytDrSinglerelationRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetRefEntId() string {
    return r._refEntId
}
// DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugKytDrSinglerelationRequest) SetDesRefEntId(_desRefEntId string) error {
    r._desRefEntId = _desRefEntId
    r.Set("des_ref_ent_id", _desRefEntId)
    return nil
}

// DesRefEntId Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetDesRefEntId() string {
    return r._desRefEntId
}
