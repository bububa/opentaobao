package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗药品召回 API请求
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

// 初始化AlibabaAlihealthDrugKytDrDrugrecalRequest对象
func NewAlibabaAlihealthDrugKytDrDrugrecalRequest() *AlibabaAlihealthDrugKytDrDrugrecalRequest{
    return &AlibabaAlihealthDrugKytDrDrugrecalRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.drugrecal"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 调用企业ID
func (r *AlibabaAlihealthDrugKytDrDrugrecalRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetRefEntId() string {
    return r.refEntId
}
// RecallBeginTime Setter
// 召回开始时间
func (r *AlibabaAlihealthDrugKytDrDrugrecalRequest) SetRecallBeginTime(recallBeginTime string) error {
    r.recallBeginTime = recallBeginTime
    r.Set("recall_begin_time", recallBeginTime)
    return nil
}

// RecallBeginTime Getter
func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetRecallBeginTime() string {
    return r.recallBeginTime
}
// RecallEndTime Setter
// 召回结束时间
func (r *AlibabaAlihealthDrugKytDrDrugrecalRequest) SetRecallEndTime(recallEndTime string) error {
    r.recallEndTime = recallEndTime
    r.Set("recall_end_time", recallEndTime)
    return nil
}

// RecallEndTime Getter
func (r AlibabaAlihealthDrugKytDrDrugrecalRequest) GetRecallEndTime() string {
    return r.recallEndTime
}
