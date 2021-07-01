package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
薄荷同步三方记录 API请求
alibaba.fmhealth.weight.lossplan.syncweightdata

用于三方薄荷同步数据到健康会员
*/
type AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest struct {
    model.Params
    // 阿里健康id
    _tpUserId   int64
    // 记录体重
    _weight   string
    // 记录日期
    _recordDate   string
}

// 初始化AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest对象
func NewAlibabaFmhealthWeightLossplanSyncweightdataRequest() *AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest{
    return &AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetApiMethodName() string {
    return "alibaba.fmhealth.weight.lossplan.syncweightdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TpUserId Setter
// 阿里健康id
func (r *AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) SetTpUserId(_tpUserId int64) error {
    r._tpUserId = _tpUserId
    r.Set("tp_user_id", _tpUserId)
    return nil
}

// TpUserId Getter
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetTpUserId() int64 {
    return r._tpUserId
}
// Weight Setter
// 记录体重
func (r *AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) SetWeight(_weight string) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetWeight() string {
    return r._weight
}
// RecordDate Setter
// 记录日期
func (r *AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) SetRecordDate(_recordDate string) error {
    r._recordDate = _recordDate
    r.Set("record_date", _recordDate)
    return nil
}

// RecordDate Getter
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetRecordDate() string {
    return r._recordDate
}
