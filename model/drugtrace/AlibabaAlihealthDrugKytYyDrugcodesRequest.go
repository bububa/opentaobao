package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品是否赋码 APIRequest
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

func NewAlibabaAlihealthDrugKytYyDrugcodesRequest() *AlibabaAlihealthDrugKytYyDrugcodesRequest{
    return &AlibabaAlihealthDrugKytYyDrugcodesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.drugcodes"
}

func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetRefEntName(refEntName string) error {
    r.refEntName = refEntName
    r.Set("ref_ent_name", refEntName)
    return nil
}

func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetRefEntName() string {
    return r.refEntName
}

func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetPhysicName(physicName string) error {
    r.physicName = physicName
    r.Set("physic_name", physicName)
    return nil
}

func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetPhysicName() string {
    return r.physicName
}

func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetProduceBatchNo(produceBatchNo string) error {
    r.produceBatchNo = produceBatchNo
    r.Set("produce_batch_no", produceBatchNo)
    return nil
}

func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetProduceBatchNo() string {
    return r.produceBatchNo
}

func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetPhysicType(physicType string) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetPhysicType() string {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetPkgSpec(pkgSpec string) error {
    r.pkgSpec = pkgSpec
    r.Set("pkg_spec", pkgSpec)
    return nil
}

func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetPkgSpec() string {
    return r.pkgSpec
}

func (r *AlibabaAlihealthDrugKytYyDrugcodesRequest) SetPrepnSpec(prepnSpec string) error {
    r.prepnSpec = prepnSpec
    r.Set("prepn_spec", prepnSpec)
    return nil
}

func (r AlibabaAlihealthDrugKytYyDrugcodesRequest) GetPrepnSpec() string {
    return r.prepnSpec
}

