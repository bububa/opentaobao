package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest 天猫服务平台商家投诉单服务商认责接口 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.admit
//
// 天猫服务平台商家投诉单服务商认责接口
type TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest struct {
	model.Params
	// 备注
	_remark string
	// 投诉单id
	_id int64
	// 认责金额，分
	_tpAdmitResponsibleAmount int64
}

// NewTmallServicecenterAnomalyrecourseHomedecorationAdmitRequest 初始化TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest对象
func NewTmallServicecenterAnomalyrecourseHomedecorationAdmitRequest() *TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest {
	return &TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) Reset() {
	r._remark = ""
	r._id = 0
	r._tpAdmitResponsibleAmount = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.admit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 备注
func (r *TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) GetRemark() string {
	return r._remark
}

// SetId is Id Setter
// 投诉单id
func (r *TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) GetId() int64 {
	return r._id
}

// SetTpAdmitResponsibleAmount is TpAdmitResponsibleAmount Setter
// 认责金额，分
func (r *TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) SetTpAdmitResponsibleAmount(_tpAdmitResponsibleAmount int64) error {
	r._tpAdmitResponsibleAmount = _tpAdmitResponsibleAmount
	r.Set("tp_admit_responsible_amount", _tpAdmitResponsibleAmount)
	return nil
}

// GetTpAdmitResponsibleAmount TpAdmitResponsibleAmount Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) GetTpAdmitResponsibleAmount() int64 {
	return r._tpAdmitResponsibleAmount
}

var poolTmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterAnomalyrecourseHomedecorationAdmitRequest()
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationAdmitRequest 从 sync.Pool 获取 TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest
func GetTmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest() *TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest {
	return poolTmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest.Get().(*TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest 将 TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest(v *TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest) {
	v.Reset()
	poolTmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest.Put(v)
}
