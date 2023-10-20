package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest 天猫服务平台服务商查询商家投诉单 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.querybyid
//
// 天猫服务平台服务商查询商家投诉单
type TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest struct {
	model.Params
	// 投诉单id
	_anomalyRecourseId int64
}

// NewTmallservicecenteranomalyrecoursehomedecorationquerybyidRequest 初始化TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest对象
func NewTmallservicecenteranomalyrecoursehomedecorationquerybyidRequest() *TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest {
	return &TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.querybyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAnomalyRecourseId is AnomalyRecourseId Setter
// 投诉单id
func (r *TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest) SetAnomalyRecourseId(_anomalyRecourseId int64) error {
	r._anomalyRecourseId = _anomalyRecourseId
	r.Set("anomaly_recourse_id", _anomalyRecourseId)
	return nil
}

// GetAnomalyRecourseId AnomalyRecourseId Getter
func (r TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest) GetAnomalyRecourseId() int64 {
	return r._anomalyRecourseId
}
