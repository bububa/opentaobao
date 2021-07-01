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
type AlibabaAlihealthDrugKytDrSinglerelationAPIRequest struct {
    model.Params
    // 追溯码
    _code   string
    // 接口调用企业的唯一标识（接口调用者）
    _refEntId   string
    // 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
    _desRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytDrSinglerelationAPIRequest对象
func NewAlibabaAlihealthDrugKytDrSinglerelationRequest() *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest{
    return &AlibabaAlihealthDrugKytDrSinglerelationAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.singlerelation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetCode() string {
    return r._code
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
    r._desRefEntId = _desRefEntId
    r.Set("des_ref_ent_id", _desRefEntId)
    return nil
}

// DesRefEntId Getter
func (r AlibabaAlihealthDrugKytDrSinglerelationAPIRequest) GetDesRefEntId() string {
    return r._desRefEntId
}
