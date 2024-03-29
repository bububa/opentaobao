package vaccin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeUserCreateAPIRequest 支付宝医疗健康疫苗用户创建 API请求
// alibaba.health.vaccin.notice.user.create
//
// 支付宝医疗健康疫苗用户创建
type AlibabaHealthVaccinNoticeUserCreateAPIRequest struct {
	model.Params
	// 支付宝用户id
	_aliPayUserId string
	// 外部渠道用户id
	_outerUserId string
	// 用户电话号码
	_mobile string
}

// NewAlibabaHealthVaccinNoticeUserCreateRequest 初始化AlibabaHealthVaccinNoticeUserCreateAPIRequest对象
func NewAlibabaHealthVaccinNoticeUserCreateRequest() *AlibabaHealthVaccinNoticeUserCreateAPIRequest {
	return &AlibabaHealthVaccinNoticeUserCreateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHealthVaccinNoticeUserCreateAPIRequest) Reset() {
	r._aliPayUserId = ""
	r._outerUserId = ""
	r._mobile = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.user.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAliPayUserId is AliPayUserId Setter
// 支付宝用户id
func (r *AlibabaHealthVaccinNoticeUserCreateAPIRequest) SetAliPayUserId(_aliPayUserId string) error {
	r._aliPayUserId = _aliPayUserId
	r.Set("ali_pay_user_id", _aliPayUserId)
	return nil
}

// GetAliPayUserId AliPayUserId Getter
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetAliPayUserId() string {
	return r._aliPayUserId
}

// SetOuterUserId is OuterUserId Setter
// 外部渠道用户id
func (r *AlibabaHealthVaccinNoticeUserCreateAPIRequest) SetOuterUserId(_outerUserId string) error {
	r._outerUserId = _outerUserId
	r.Set("outer_user_id", _outerUserId)
	return nil
}

// GetOuterUserId OuterUserId Getter
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetOuterUserId() string {
	return r._outerUserId
}

// SetMobile is Mobile Setter
// 用户电话号码
func (r *AlibabaHealthVaccinNoticeUserCreateAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaHealthVaccinNoticeUserCreateAPIRequest) GetMobile() string {
	return r._mobile
}

var poolAlibabaHealthVaccinNoticeUserCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHealthVaccinNoticeUserCreateRequest()
	},
}

// GetAlibabaHealthVaccinNoticeUserCreateRequest 从 sync.Pool 获取 AlibabaHealthVaccinNoticeUserCreateAPIRequest
func GetAlibabaHealthVaccinNoticeUserCreateAPIRequest() *AlibabaHealthVaccinNoticeUserCreateAPIRequest {
	return poolAlibabaHealthVaccinNoticeUserCreateAPIRequest.Get().(*AlibabaHealthVaccinNoticeUserCreateAPIRequest)
}

// ReleaseAlibabaHealthVaccinNoticeUserCreateAPIRequest 将 AlibabaHealthVaccinNoticeUserCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeUserCreateAPIRequest(v *AlibabaHealthVaccinNoticeUserCreateAPIRequest) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeUserCreateAPIRequest.Put(v)
}
