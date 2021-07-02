package tmallservice

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerTaglistGetAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.taglist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerTaglistGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
