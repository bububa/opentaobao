package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗药品召回 APIRequest
alibaba.alihealth.drug.kyt.dr.drugrecal

生产企业发布的召回信息，按照批次进行召回，收货和发货环节的单据处理中调用接口进行查询；
*/
type AlibabaAlihealthDrugKytDrDrugrecalRequest struct {
    model.Params

    // 调用企业ID
    refEntId   string 

    // 召回开始时间
    recallBeginTime   string 

    // 召回结束时间
    recallEndTime   string 

}

func NewAlibabaAlihealthDrugKytDrDrugrecalRequest() *AlibabaAlihealthDrugKytDrDrugrecalRequest{
    return &AlibabaAlihealthDrugKytDrDrugrecalRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.drugrecal"
}

func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrDrugrecalRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrDrugrecalRequest) SetRecallBeginTime(recallBeginTime string) error {
    r.recallBeginTime = recallBeginTime
    r.Set("recall_begin_time", recallBeginTime)
    return nil
}

func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetRecallBeginTime() string {
    return r.recallBeginTime
}

func (r *AlibabaAlihealthDrugKytDrDrugrecalRequest) SetRecallEndTime(recallEndTime string) error {
    r.recallEndTime = recallEndTime
    r.Set("recall_end_time", recallEndTime)
    return nil
}

func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetRecallEndTime() string {
    return r.recallEndTime
}

