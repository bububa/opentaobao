package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多融通过企业ID得到一个企业的详细信息 APIRequest
alibaba.alihealth.drug.kyt.dr.getbyentid

根据企业主键查看企业详细信息
*/
type AlibabaAlihealthDrugKytDrGetbyentidRequest struct {
    model.Params

    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string 

    // 准备要查询的企业ID（返回该企业ID的详细信息）
    entId   string 

}

func NewAlibabaAlihealthDrugKytDrGetbyentidRequest() *AlibabaAlihealthDrugKytDrGetbyentidRequest{
    return &AlibabaAlihealthDrugKytDrGetbyentidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrGetbyentidRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.getbyentid"
}

func (r AlibabaAlihealthDrugKytDrGetbyentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrGetbyentidRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrGetbyentidRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrGetbyentidRequest) SetEntId(entId string) error {
    r.entId = entId
    r.Set("ent_id", entId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrGetbyentidRequest) GetEntId() string {
    return r.entId
}

