package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseCloseAPIRequest 服务投诉问题单关单 API请求
// tmall.servicecenter.anomalyrecourse.close
//
// 提供给服务商在投诉单完结时调用，关闭投诉问题工单。
type TmallServicecenterAnomalyrecourseCloseAPIRequest struct {
	model.Params
	// 完结时服务商自定义回复给消费者内容
	_remark string
	// 投诉单号
	_id int64
}

// NewTmallServicecenterAnomalyrecourseCloseRequest 初始化TmallServicecenterAnomalyrecourseCloseAPIRequest对象
func NewTmallServicecenterAnomalyrecourseCloseRequest() *TmallServicecenterAnomalyrecourseCloseAPIRequest {
	return &TmallServicecenterAnomalyrecourseCloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseCloseAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseCloseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRemark is Remark Setter
// 完结时服务商自定义回复给消费者内容
func (r *TmallServicecenterAnomalyrecourseCloseAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallServicecenterAnomalyrecourseCloseAPIRequest) GetRemark() string {
	return r._remark
}

// SetId is Id Setter
// 投诉单号
func (r *TmallServicecenterAnomalyrecourseCloseAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterAnomalyrecourseCloseAPIRequest) GetId() int64 {
	return r._id
}
