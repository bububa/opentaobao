package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrDrugrecalAPIRequest 疫苗药品召回 API请求
// alibaba.alihealth.drug.kyt.dr.drugrecal
//
// 生产企业发布的召回信息，按照批次进行召回，收货和发货环节的单据处理中调用接口进行查询；
type AlibabaAlihealthDrugKytDrDrugrecalAPIRequest struct {
	model.Params
	// 调用企业ID
	_refEntId string
	// 召回开始时间
	_recallBeginTime string
	// 召回结束时间
	_recallEndTime string
}

// NewAlibabaAlihealthDrugKytDrDrugrecalRequest 初始化AlibabaAlihealthDrugKytDrDrugrecalAPIRequest对象
func NewAlibabaAlihealthDrugKytDrDrugrecalRequest() *AlibabaAlihealthDrugKytDrDrugrecalAPIRequest {
	return &AlibabaAlihealthDrugKytDrDrugrecalAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) Reset() {
	r._refEntId = ""
	r._recallBeginTime = ""
	r._recallEndTime = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.drugrecal"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用企业ID
func (r *AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetRecallBeginTime is RecallBeginTime Setter
// 召回开始时间
func (r *AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) SetRecallBeginTime(_recallBeginTime string) error {
	r._recallBeginTime = _recallBeginTime
	r.Set("recall_begin_time", _recallBeginTime)
	return nil
}

// GetRecallBeginTime RecallBeginTime Getter
func (r AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) GetRecallBeginTime() string {
	return r._recallBeginTime
}

// SetRecallEndTime is RecallEndTime Setter
// 召回结束时间
func (r *AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) SetRecallEndTime(_recallEndTime string) error {
	r._recallEndTime = _recallEndTime
	r.Set("recall_end_time", _recallEndTime)
	return nil
}

// GetRecallEndTime RecallEndTime Getter
func (r AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) GetRecallEndTime() string {
	return r._recallEndTime
}

var poolAlibabaAlihealthDrugKytDrDrugrecalAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrDrugrecalRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrDrugrecalRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrDrugrecalAPIRequest
func GetAlibabaAlihealthDrugKytDrDrugrecalAPIRequest() *AlibabaAlihealthDrugKytDrDrugrecalAPIRequest {
	return poolAlibabaAlihealthDrugKytDrDrugrecalAPIRequest.Get().(*AlibabaAlihealthDrugKytDrDrugrecalAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrDrugrecalAPIRequest 将 AlibabaAlihealthDrugKytDrDrugrecalAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrDrugrecalAPIRequest(v *AlibabaAlihealthDrugKytDrDrugrecalAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrDrugrecalAPIRequest.Put(v)
}
