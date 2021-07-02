package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest 查询外部小件员休息 API请求
// cainiao.nborderfront.user.outside.queryoutsideuser
//
// 采用SPI方式查询外部公司的小件员信息
type CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest struct {
	model.Params
	// cpcode
	_cpCode string
	// cp小件员ID
	_cpUserId string
}

// NewCainiaoNborderfrontUserOutsideQueryoutsideuserRequest 初始化CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest对象
func NewCainiaoNborderfrontUserOutsideQueryoutsideuserRequest() *CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest {
	return &CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest) GetApiMethodName() string {
	return "cainiao.nborderfront.user.outside.queryoutsideuser"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CpCode Setter
// cpcode
func (r *CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// Get CpCode Getter
func (r CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest) GetCpCode() string {
	return r._cpCode
}

// Set is CpUserId Setter
// cp小件员ID
func (r *CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest) SetCpUserId(_cpUserId string) error {
	r._cpUserId = _cpUserId
	r.Set("cp_user_id", _cpUserId)
	return nil
}

// Get CpUserId Getter
func (r CainiaoNborderfrontUserOutsideQueryoutsideuserAPIRequest) GetCpUserId() string {
	return r._cpUserId
}
