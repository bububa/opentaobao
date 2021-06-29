package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
薄荷同步三方记录 APIRequest
alibaba.fmhealth.weight.lossplan.syncweightdata

用于三方薄荷同步数据到健康会员
*/
type AlibabaFmhealthWeightLossplanSyncweightdataRequest struct {
    model.Params

    // 阿里健康id
    tpUserId   int64 

    // 记录体重
    weight   string 

    // 记录日期
    recordDate   string 

}

func NewAlibabaFmhealthWeightLossplanSyncweightdataRequest() *AlibabaFmhealthWeightLossplanSyncweightdataRequest{
    return &AlibabaFmhealthWeightLossplanSyncweightdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFmhealthWeightLossplanSyncweightdataRequest) GetApiMethodName() string {
    return "alibaba.fmhealth.weight.lossplan.syncweightdata"
}

func (r AlibabaFmhealthWeightLossplanSyncweightdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFmhealthWeightLossplanSyncweightdataRequest) SetTpUserId(tpUserId int64) error {
    r.tpUserId = tpUserId
    r.Set("tp_user_id", tpUserId)
    return nil
}

func (r AlibabaFmhealthWeightLossplanSyncweightdataRequest) GetTpUserId() int64 {
    return r.tpUserId
}

func (r *AlibabaFmhealthWeightLossplanSyncweightdataRequest) SetWeight(weight string) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r AlibabaFmhealthWeightLossplanSyncweightdataRequest) GetWeight() string {
    return r.weight
}

func (r *AlibabaFmhealthWeightLossplanSyncweightdataRequest) SetRecordDate(recordDate string) error {
    r.recordDate = recordDate
    r.Set("record_date", recordDate)
    return nil
}

func (r AlibabaFmhealthWeightLossplanSyncweightdataRequest) GetRecordDate() string {
    return r.recordDate
}

