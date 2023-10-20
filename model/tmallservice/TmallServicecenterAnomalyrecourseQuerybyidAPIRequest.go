package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseQuerybyidAPIRequest 根据一键求助id查询指定服务商的一键求助单 API请求
// tmall.servicecenter.anomalyrecourse.querybyid
//
// 根据一键求助id查询指定服务商的一键求助单
type TmallServicecenterAnomalyrecourseQuerybyidAPIRequest struct {
	model.Params
	// 一键求助的id
	_anomalyRecourseId int64
}

// NewTmallServicecenterAnomalyrecourseQuerybyidRequest 初始化TmallServicecenterAnomalyrecourseQuerybyidAPIRequest对象
func NewTmallServicecenterAnomalyrecourseQuerybyidRequest() *TmallServicecenterAnomalyrecourseQuerybyidAPIRequest {
	return &TmallServicecenterAnomalyrecourseQuerybyidAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) Reset() {
	r._anomalyRecourseId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.querybyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAnomalyRecourseId is AnomalyRecourseId Setter
// 一键求助的id
func (r *TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) SetAnomalyRecourseId(_anomalyRecourseId int64) error {
	r._anomalyRecourseId = _anomalyRecourseId
	r.Set("anomaly_recourse_id", _anomalyRecourseId)
	return nil
}

// GetAnomalyRecourseId AnomalyRecourseId Getter
func (r TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) GetAnomalyRecourseId() int64 {
	return r._anomalyRecourseId
}

var poolTmallServicecenterAnomalyrecourseQuerybyidAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterAnomalyrecourseQuerybyidRequest()
	},
}

// GetTmallServicecenterAnomalyrecourseQuerybyidRequest 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseQuerybyidAPIRequest
func GetTmallServicecenterAnomalyrecourseQuerybyidAPIRequest() *TmallServicecenterAnomalyrecourseQuerybyidAPIRequest {
	return poolTmallServicecenterAnomalyrecourseQuerybyidAPIRequest.Get().(*TmallServicecenterAnomalyrecourseQuerybyidAPIRequest)
}

// ReleaseTmallServicecenterAnomalyrecourseQuerybyidAPIRequest 将 TmallServicecenterAnomalyrecourseQuerybyidAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseQuerybyidAPIRequest(v *TmallServicecenterAnomalyrecourseQuerybyidAPIRequest) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseQuerybyidAPIRequest.Put(v)
}
