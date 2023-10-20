package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerTaglistGetAPIRequest 获取工人标签 API请求
// tmall.servicecenter.worker.taglist.get
//
// 服务商获取对应工人的标签
type TmallServicecenterWorkerTaglistGetAPIRequest struct {
	model.Params
	// 工人注册勤鸽时的身份证
	_idNumber string
	// 工人注册勤鸽时的手机号码
	_mobile string
}

// NewTmallServicecenterWorkerTaglistGetRequest 初始化TmallServicecenterWorkerTaglistGetAPIRequest对象
func NewTmallServicecenterWorkerTaglistGetRequest() *TmallServicecenterWorkerTaglistGetAPIRequest {
	return &TmallServicecenterWorkerTaglistGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkerTaglistGetAPIRequest) Reset() {
	r._idNumber = ""
	r._mobile = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerTaglistGetAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.taglist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerTaglistGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkerTaglistGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdNumber is IdNumber Setter
// 工人注册勤鸽时的身份证
func (r *TmallServicecenterWorkerTaglistGetAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r TmallServicecenterWorkerTaglistGetAPIRequest) GetIdNumber() string {
	return r._idNumber
}

// SetMobile is Mobile Setter
// 工人注册勤鸽时的手机号码
func (r *TmallServicecenterWorkerTaglistGetAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TmallServicecenterWorkerTaglistGetAPIRequest) GetMobile() string {
	return r._mobile
}

var poolTmallServicecenterWorkerTaglistGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkerTaglistGetRequest()
	},
}

// GetTmallServicecenterWorkerTaglistGetRequest 从 sync.Pool 获取 TmallServicecenterWorkerTaglistGetAPIRequest
func GetTmallServicecenterWorkerTaglistGetAPIRequest() *TmallServicecenterWorkerTaglistGetAPIRequest {
	return poolTmallServicecenterWorkerTaglistGetAPIRequest.Get().(*TmallServicecenterWorkerTaglistGetAPIRequest)
}

// ReleaseTmallServicecenterWorkerTaglistGetAPIRequest 将 TmallServicecenterWorkerTaglistGetAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkerTaglistGetAPIRequest(v *TmallServicecenterWorkerTaglistGetAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkerTaglistGetAPIRequest.Put(v)
}
