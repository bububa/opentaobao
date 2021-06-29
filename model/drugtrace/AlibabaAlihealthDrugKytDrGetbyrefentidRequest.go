package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多融通过一个企业唯一标识查询企业详细信息 APIRequest
alibaba.alihealth.drug.kyt.dr.getbyrefentid

根据企业唯一标识查看企业详细信息
*/
type AlibabaAlihealthDrugKytDrGetbyrefentidRequest struct {
    model.Params

    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string 

    // 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
    destRefEntId   string 

}

func NewAlibabaAlihealthDrugKytDrGetbyrefentidRequest() *AlibabaAlihealthDrugKytDrGetbyrefentidRequest{
    return &AlibabaAlihealthDrugKytDrGetbyrefentidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrGetbyrefentidRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.getbyrefentid"
}

func (r AlibabaAlihealthDrugKytDrGetbyrefentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrGetbyrefentidRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrGetbyrefentidRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrGetbyrefentidRequest) SetDestRefEntId(destRefEntId string) error {
    r.destRefEntId = destRefEntId
    r.Set("dest_ref_ent_id", destRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrGetbyrefentidRequest) GetDestRefEntId() string {
    return r.destRefEntId
}

