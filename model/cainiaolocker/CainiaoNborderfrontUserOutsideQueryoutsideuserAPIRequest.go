package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest 查询外部小件员休息 API请求
// cainiao.nborderfront.user.outside.queryoutsideuser
//
// 采用SPI方式查询外部公司的小件员信息
type CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest struct {
	model.Params
	// cpcode
	_cpCode string
	// cp小件员ID
	_cpUserId string
}

// NewCainiaonborderfrontuseroutsidequeryoutsideuserRequest 初始化CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest对象
func NewCainiaonborderfrontuseroutsidequeryoutsideuserRequest() *CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest {
	return &CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest) GetApiMethodName() string {
	return "cainiao.nborderfront.user.outside.queryoutsideuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// cpcode
func (r *CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest) GetCpCode() string {
	return r._cpCode
}

// SetCpUserId is CpUserId Setter
// cp小件员ID
func (r *CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest) SetCpUserId(_cpUserId string) error {
	r._cpUserId = _cpUserId
	r.Set("cp_user_id", _cpUserId)
	return nil
}

// GetCpUserId CpUserId Getter
func (r CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest) GetCpUserId() string {
	return r._cpUserId
}
