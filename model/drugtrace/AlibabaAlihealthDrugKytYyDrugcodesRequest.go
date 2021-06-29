package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品是否赋码 API请求
alibaba.alihealth.drug.kyt.yy.drugcodes

药品是否赋码
*/
type AlibabaAlihealthDrugKytYyDrugcodesRequest struct {
    model.Params
    // 企业名称
    refEntName   string
    // 药品名称
    physicName   string
    // 生产批号
    produceBatchNo   string
    // 药品类型
    physicType   string
    // 包装规格
    pkgSpec   string
    // 制剂规格
    prepnSpec   string
}

// 初始化AlibabaAlihealthDrugKytYyDrugcodesRequest对象
func NewAlibabaAlihealthDrugKytYyDrugcodesRequest() *AlibabaAlihealthDrugKytYyDrugcodesRequest{
    return &AlibabaAlihealthDrugKytYyDrugcodesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.drugcodes"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetRefEntName(refEntName string) error {
    r.refEntName = refEntName
    r.Set("ref_ent_name", refEntName)
    return nil
}

// RefEntName Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetRefEntName() string {
    return r.refEntName
}
// PhysicName Setter
// 药品名称
func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetPhysicName(physicName string) error {
    r.physicName = physicName
    r.Set("physic_name", physicName)
    return nil
}

// PhysicName Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetPhysicName() string {
    return r.physicName
}
// ProduceBatchNo Setter
// 生产批号
func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetProduceBatchNo(produceBatchNo string) error {
    r.produceBatchNo = produceBatchNo
    r.Set("produce_batch_no", produceBatchNo)
    return nil
}

// ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetProduceBatchNo() string {
    return r.produceBatchNo
}
// PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetPhysicType(physicType string) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetPhysicType() string {
    return r.physicType
}
// PkgSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetPkgSpec(pkgSpec string) error {
    r.pkgSpec = pkgSpec
    r.Set("pkg_spec", pkgSpec)
    return nil
}

// PkgSpec Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetPkgSpec() string {
    return r.pkgSpec
}
// PrepnSpec Setter
// 制剂规格
func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetPrepnSpec(prepnSpec string) error {
    r.prepnSpec = prepnSpec
    r.Set("prepn_spec", prepnSpec)
    return nil
}

// PrepnSpec Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetPrepnSpec() string {
    return r.prepnSpec
}
