package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
药品是否赋码 APIRequest
alibaba.alihealth.drug.kyt.drugcodes

药品是否赋码
*/
type AlibabaAlihealthDrugKytDrugcodesRequest struct {
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

func NewAlibabaAlihealthDrugKytDrugcodesRequest() *AlibabaAlihealthDrugKytDrugcodesRequest{
    return &AlibabaAlihealthDrugKytDrugcodesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrugcodesRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.drugcodes"
}

func (r AlibabaAlihealthDrugKytDrugcodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrugcodesRequest) SetRefEntName(refEntName string) error {
    r.refEntName = refEntName
    r.Set("ref_ent_name", refEntName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugcodesRequest) GetRefEntName() string {
    return r.refEntName
}

func (r *AlibabaAlihealthDrugKytDrugcodesRequest) SetPhysicName(physicName string) error {
    r.physicName = physicName
    r.Set("physic_name", physicName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugcodesRequest) GetPhysicName() string {
    return r.physicName
}

func (r *AlibabaAlihealthDrugKytDrugcodesRequest) SetProduceBatchNo(produceBatchNo string) error {
    r.produceBatchNo = produceBatchNo
    r.Set("produce_batch_no", produceBatchNo)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugcodesRequest) GetProduceBatchNo() string {
    return r.produceBatchNo
}

func (r *AlibabaAlihealthDrugKytDrugcodesRequest) SetPhysicType(physicType string) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugcodesRequest) GetPhysicType() string {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytDrugcodesRequest) SetPkgSpec(pkgSpec string) error {
    r.pkgSpec = pkgSpec
    r.Set("pkg_spec", pkgSpec)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugcodesRequest) GetPkgSpec() string {
    return r.pkgSpec
}

func (r *AlibabaAlihealthDrugKytDrugcodesRequest) SetPrepnSpec(prepnSpec string) error {
    r.prepnSpec = prepnSpec
    r.Set("prepn_spec", prepnSpec)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugcodesRequest) GetPrepnSpec() string {
    return r.prepnSpec
}

