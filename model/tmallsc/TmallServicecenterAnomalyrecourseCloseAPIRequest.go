package tmallsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterAnomalyrecourseCloseAPIRequest) Reset() {
	r._remark = ""
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseCloseAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterAnomalyrecourseCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallServicecenterAnomalyrecourseCloseAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterAnomalyrecourseCloseRequest()
	},
}

// GetTmallServicecenterAnomalyrecourseCloseRequest 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseCloseAPIRequest
func GetTmallServicecenterAnomalyrecourseCloseAPIRequest() *TmallServicecenterAnomalyrecourseCloseAPIRequest {
	return poolTmallServicecenterAnomalyrecourseCloseAPIRequest.Get().(*TmallServicecenterAnomalyrecourseCloseAPIRequest)
}

// ReleaseTmallServicecenterAnomalyrecourseCloseAPIRequest 将 TmallServicecenterAnomalyrecourseCloseAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseCloseAPIRequest(v *TmallServicecenterAnomalyrecourseCloseAPIRequest) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseCloseAPIRequest.Put(v)
}
