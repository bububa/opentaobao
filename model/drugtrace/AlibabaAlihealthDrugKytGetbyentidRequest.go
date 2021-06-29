package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据企业主键查看企业详细信息 APIRequest
alibaba.alihealth.drug.kyt.getbyentid

根据企业主键查看企业详细信息
*/
type AlibabaAlihealthDrugKytGetbyentidRequest struct {
    model.Params

    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string 

    // 准备要查询的企业ID（返回该企业ID的详细信息）
    entId   string 

}

func NewAlibabaAlihealthDrugKytGetbyentidRequest() *AlibabaAlihealthDrugKytGetbyentidRequest{
    return &AlibabaAlihealthDrugKytGetbyentidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytGetbyentidRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getbyentid"
}

func (r AlibabaAlihealthDrugKytGetbyentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytGetbyentidRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetbyentidRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytGetbyentidRequest) SetEntId(entId string) error {
    r.entId = entId
    r.Set("ent_id", entId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetbyentidRequest) GetEntId() string {
    return r.entId
}

