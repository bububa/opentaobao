package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据企业唯一标识查看企业详细信息 APIRequest
alibaba.alihealth.drug.kyt.getbyrefentid

根据企业唯一标识查看企业详细信息
*/
type AlibabaAlihealthDrugKytGetbyrefentidRequest struct {
    model.Params

    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string 

    // 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
    destRefEntId   string 

}

func NewAlibabaAlihealthDrugKytGetbyrefentidRequest() *AlibabaAlihealthDrugKytGetbyrefentidRequest{
    return &AlibabaAlihealthDrugKytGetbyrefentidRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytGetbyrefentidRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getbyrefentid"
}

func (r AlibabaAlihealthDrugKytGetbyrefentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytGetbyrefentidRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetbyrefentidRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytGetbyrefentidRequest) SetDestRefEntId(destRefEntId string) error {
    r.destRefEntId = destRefEntId
    r.Set("dest_ref_ent_id", destRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetbyrefentidRequest) GetDestRefEntId() string {
    return r.destRefEntId
}

