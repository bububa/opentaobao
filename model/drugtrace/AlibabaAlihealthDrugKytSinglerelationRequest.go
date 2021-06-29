package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单码关联关系查询，通过一个码查询这个码下的所有子码 API请求
alibaba.alihealth.drug.kyt.singlerelation

单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
*/
type AlibabaAlihealthDrugKytSinglerelationRequest struct {
    model.Params
    // 追溯码
    code   string
    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string
    // 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
    desRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytSinglerelationRequest对象
func NewAlibabaAlihealthDrugKytSinglerelationRequest() *AlibabaAlihealthDrugKytSinglerelationRequest{
    return &AlibabaAlihealthDrugKytSinglerelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSinglerelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.singlerelation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSinglerelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytSinglerelationRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugKytSinglerelationRequest) GetCode() string {
    return r.code
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytSinglerelationRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytSinglerelationRequest) GetRefEntId() string {
    return r.refEntId
}
// DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugKytSinglerelationRequest) SetDesRefEntId(desRefEntId string) error {
    r.desRefEntId = desRefEntId
    r.Set("des_ref_ent_id", desRefEntId)
    return nil
}

// DesRefEntId Getter
func (r AlibabaAlihealthDrugKytSinglerelationRequest) GetDesRefEntId() string {
    return r.desRefEntId
}
