package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinnoticeusercreateAPIRequest 支付宝医疗健康疫苗用户创建 API请求
// alibaba.health.vaccin.notice.user.create
//
// 支付宝医疗健康疫苗用户创建
type AlibabahealthvaccinnoticeusercreateAPIRequest struct {
	model.Params
	// 支付宝用户id
	_aliPayUserId string
	// 外部渠道用户id
	_outerUserId string
	// 用户电话号码
	_mobile string
}

// NewAlibabahealthvaccinnoticeusercreateRequest 初始化AlibabahealthvaccinnoticeusercreateAPIRequest对象
func NewAlibabahealthvaccinnoticeusercreateRequest() *AlibabahealthvaccinnoticeusercreateAPIRequest {
	return &AlibabahealthvaccinnoticeusercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinnoticeusercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.user.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinnoticeusercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinnoticeusercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAliPayUserId is AliPayUserId Setter
// 支付宝用户id
func (r *AlibabahealthvaccinnoticeusercreateAPIRequest) SetAliPayUserId(_aliPayUserId string) error {
	r._aliPayUserId = _aliPayUserId
	r.Set("ali_pay_user_id", _aliPayUserId)
	return nil
}

// GetAliPayUserId AliPayUserId Getter
func (r AlibabahealthvaccinnoticeusercreateAPIRequest) GetAliPayUserId() string {
	return r._aliPayUserId
}

// SetOuterUserId is OuterUserId Setter
// 外部渠道用户id
func (r *AlibabahealthvaccinnoticeusercreateAPIRequest) SetOuterUserId(_outerUserId string) error {
	r._outerUserId = _outerUserId
	r.Set("outer_user_id", _outerUserId)
	return nil
}

// GetOuterUserId OuterUserId Getter
func (r AlibabahealthvaccinnoticeusercreateAPIRequest) GetOuterUserId() string {
	return r._outerUserId
}

// SetMobile is Mobile Setter
// 用户电话号码
func (r *AlibabahealthvaccinnoticeusercreateAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabahealthvaccinnoticeusercreateAPIRequest) GetMobile() string {
	return r._mobile
}
