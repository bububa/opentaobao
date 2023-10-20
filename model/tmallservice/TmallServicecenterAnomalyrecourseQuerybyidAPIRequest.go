package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursequerybyidAPIRequest 根据一键求助id查询指定服务商的一键求助单 API请求
// tmall.servicecenter.anomalyrecourse.querybyid
//
// 根据一键求助id查询指定服务商的一键求助单
type TmallservicecenteranomalyrecoursequerybyidAPIRequest struct {
	model.Params
	// 一键求助的id
	_anomalyRecourseId int64
}

// NewTmallservicecenteranomalyrecoursequerybyidRequest 初始化TmallservicecenteranomalyrecoursequerybyidAPIRequest对象
func NewTmallservicecenteranomalyrecoursequerybyidRequest() *TmallservicecenteranomalyrecoursequerybyidAPIRequest {
	return &TmallservicecenteranomalyrecoursequerybyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursequerybyidAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.querybyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursequerybyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursequerybyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAnomalyRecourseId is AnomalyRecourseId Setter
// 一键求助的id
func (r *TmallservicecenteranomalyrecoursequerybyidAPIRequest) SetAnomalyRecourseId(_anomalyRecourseId int64) error {
	r._anomalyRecourseId = _anomalyRecourseId
	r.Set("anomaly_recourse_id", _anomalyRecourseId)
	return nil
}

// GetAnomalyRecourseId AnomalyRecourseId Getter
func (r TmallservicecenteranomalyrecoursequerybyidAPIRequest) GetAnomalyRecourseId() int64 {
	return r._anomalyRecourseId
}
