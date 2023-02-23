package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest 薄荷同步三方记录 API请求
// alibaba.fmhealth.weight.lossplan.syncweightdata
//
// 用于三方薄荷同步数据到健康会员
type AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest struct {
	model.Params
	// 记录体重
	_weight string
	// 记录日期
	_recordDate string
	// 阿里健康id
	_tpUserId int64
}

// NewAlibabaFmhealthWeightLossplanSyncweightdataRequest 初始化AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest对象
func NewAlibabaFmhealthWeightLossplanSyncweightdataRequest() *AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest {
	return &AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetApiMethodName() string {
	return "alibaba.fmhealth.weight.lossplan.syncweightdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWeight is Weight Setter
// 记录体重
func (r *AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) SetWeight(_weight string) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetWeight() string {
	return r._weight
}

// SetRecordDate is RecordDate Setter
// 记录日期
func (r *AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) SetRecordDate(_recordDate string) error {
	r._recordDate = _recordDate
	r.Set("record_date", _recordDate)
	return nil
}

// GetRecordDate RecordDate Getter
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetRecordDate() string {
	return r._recordDate
}

// SetTpUserId is TpUserId Setter
// 阿里健康id
func (r *AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) SetTpUserId(_tpUserId int64) error {
	r._tpUserId = _tpUserId
	r.Set("tp_user_id", _tpUserId)
	return nil
}

// GetTpUserId TpUserId Getter
func (r AlibabaFmhealthWeightLossplanSyncweightdataAPIRequest) GetTpUserId() int64 {
	return r._tpUserId
}
